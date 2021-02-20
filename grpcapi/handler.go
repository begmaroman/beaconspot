package grpcapi

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"go.uber.org/zap"

	"github.com/begmaroman/beaconspot/beaconchain"
	grpcapiproto "github.com/begmaroman/beaconspot/proto/beaconspot"
	"github.com/begmaroman/beaconspot/proto/health"
)

// To make sure Handler implements grpcapiproto.BeaconSpotServiceHandler interface.
var _ grpcapiproto.BeaconSpotServiceHandler = &handler{}

// Options serves as the dependency injection container to create a new handler.
type Options struct {
	BeaconChainClient beaconchain.BeaconChain
	SelfPingClient    *health.SelfPingClient
	Log               *zap.Logger
}

// handler implements grpcapiproto.BeaconSpotServiceHandler interface
type handler struct {
	beaconChainClient beaconchain.BeaconChain
	selfPingClient    *health.SelfPingClient
	log               *zap.Logger
}

// New is the constructor of handler
func New(opts Options) grpcapiproto.BeaconSpotServiceHandler {
	return &handler{
		beaconChainClient: opts.BeaconChainClient,
		selfPingClient:    opts.SelfPingClient,
		log:               opts.Log,
	}
}

func (h *handler) Health(context.Context, *empty.Empty, *health.HealthResponse) error {
	// TODO: Implement healthcheck
	return nil
}

func (h *handler) Ping(context.Context, *empty.Empty, *empty.Empty) error {
	return nil
}
