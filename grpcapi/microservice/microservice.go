package microservice

import (
	"net"

	"github.com/bloxapp/eth2-key-manager/core"
	"github.com/pkg/errors"
	ethpb "github.com/prysmaticlabs/prysm/proto/prysm/v1alpha1"
	"google.golang.org/grpc"

	"github.com/begmaroman/beaconspot/beaconchain"
	"github.com/begmaroman/beaconspot/beaconchain/combined"
	"github.com/begmaroman/beaconspot/beaconchain/lighthouse"
	"github.com/begmaroman/beaconspot/beaconchain/prysm"
	"github.com/begmaroman/beaconspot/grpcapi"
	beaconspotproto "github.com/begmaroman/beaconspot/proto/beaconspot"
	"github.com/begmaroman/beaconspot/utils/grpcex"
)

var opts Options

// MicroService is the micro-service.
type MicroService struct {
	grpcServer *grpc.Server
	lis        net.Listener
}

// Init initializes the service.
func Init(clientOpts *ClientOptions) (*MicroService, error) {
	if err := opts.Load(); err != nil {
		return nil, err
	}

	// Validate options.
	if err := opts.Validate(); err != nil {
		return nil, errors.Wrap(err, "options validation failed")
	}

	if opts.IsTest {
		clientOpts.Log.Info("Running in test mode!")
	}

	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		return nil, err
	}
	grpcServer := grpc.NewServer()

	// Create beacon chain client
	var clients []beaconchain.BeaconChain

	// Populate Prysm clients
	for _, addr := range opts.PrysmAddrs {
		conn, err := grpcex.Dial(clientOpts.Log, addr)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to create gRPC connection with beacon chain server '%s'", addr)
		}

		beaconNodeClient := ethpb.NewBeaconNodeValidatorClient(conn)
		nodeClient := ethpb.NewNodeClient(conn)
		prysmClient := prysm.New(clientOpts.Log, beaconNodeClient, nodeClient)
		clients = append(clients, prysmClient)
	}

	// Populate Lighthouse clients
	for _, addr := range opts.LighthouseAddrs {
		if addr == "" {
			continue
		}

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
		Log:               clientOpts.Log,
	})

	// Register the service.
	beaconspotproto.RegisterBeaconSpotServiceServer(grpcServer, handler)

	return &MicroService{
		lis:        lis,
		grpcServer: grpcServer,
	}, nil
}

// Run runs the service.
func (s *MicroService) Run() error {
	return s.grpcServer.Serve(s.lis)
}
