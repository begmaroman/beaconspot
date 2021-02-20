package microservice

import (
	"github.com/begmaroman/beaconspot/beaconchain"
	"github.com/begmaroman/beaconspot/beaconchain/combined"
	"github.com/begmaroman/beaconspot/beaconchain/lighthouse"
	"github.com/begmaroman/beaconspot/beaconchain/prysm"
	"github.com/begmaroman/beaconspot/utils/grpcex"
	"github.com/bloxapp/eth2-key-manager/core"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/pkg/errors"
	ethpb "github.com/prysmaticlabs/ethereumapis/eth/v1alpha1"

	"github.com/begmaroman/beaconspot/grpcapi"
	grpcapiproto "github.com/begmaroman/beaconspot/proto/grpcapi"
	"github.com/begmaroman/beaconspot/proto/health"
	"github.com/begmaroman/beaconspot/utils/healthchecker"
)

// MicroService is the micro-service.
type MicroService struct {
	svc     micro.Service
	handler grpcapiproto.BeaconSpotServiceHandler
}

// Init initializes the service.
func Init(clientOpts *ClientOptions) (*MicroService, error) {
	// Create micro-service.
	svc := micro.NewService(
		micro.Name(clientOpts.Name),
		micro.Version(clientOpts.Version),
		micro.Flags(flags...),
		micro.Address(":5000"),
		micro.Action(func(ctx *cli.Context) error {
			return opts.Load(ctx)
		}),
		micro.BeforeStart(func() error {
			// Validate options.
			if err := opts.Validate(); err != nil {
				return errors.Wrap(err, "options validation failed")
			}

			if opts.IsTest {
				clientOpts.Log.Info("Running in test mode!")
			}

			return nil
		}),
	)

	// Parse command-line arguments.
	svc.Init()

	return New(svc, clientOpts)
}

// New is the constructor of the service.
func New(svc micro.Service, clientOpts *ClientOptions) (*MicroService, error) {
	// Create a self-pinger client.
	selfPingClient := health.NewSelfPingClient(svc, grpcapiproto.NewBeaconSpotService(clientOpts.Name, svc.Client()))

	var clients []beaconchain.BeaconChain

	// Populate Prysm clients
	for _, addr := range opts.PrysmAddrs {
		conn, err := grpcex.Dial(clientOpts.Log, addr)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to create gRPC connection with beacon chain server '%s'", addr)
		}

		beaconNodeClient := ethpb.NewBeaconNodeValidatorClient(conn)
		prysmClient := prysm.New(clientOpts.Log, beaconNodeClient)
		clients = append(clients, prysmClient)
	}

	// Populate Lighthouse clients
	for _, addr := range opts.PrysmAddrs {
		// TODO: Pass network through evars
		lightHouseClient := lighthouse.New(clientOpts.Log, core.PyrmontNetwork, addr)
		clients = append(clients, lightHouseClient)
	}

	if len(clients) == 0 {
		return nil, errors.New("no clients provided")
	}

	beaconChainClient := combined.New(clients...)

	// Create RPC handler.
	handler := grpcapi.New(grpcapi.Options{
		BeaconChainClient: beaconChainClient,
		SelfPingClient:    selfPingClient,
		Log:               clientOpts.Log,
	})

	// Register the service.
	if err := grpcapiproto.RegisterBeaconSpotServiceHandler(svc.Server(), handler); err != nil {
		return nil, errors.Wrap(err, "failed to register RPC handler")
	}

	return &MicroService{
		svc:     svc,
		handler: handler,
	}, nil
}

// Run runs the service.
func (s *MicroService) Run() error {
	// Run helathcheck endpoint.
	shutdown := healthchecker.Run(healthchecker.WrapRPC(s.handler.Health), nil)

	// Stop helathcheck endpoint after RPC service stop.
	s.svc.Init(micro.AfterStop(shutdown))

	// Start service.
	if err := s.svc.Run(); err != nil {
		return errors.Wrap(err, "failed to run")
	}

	return nil
}
