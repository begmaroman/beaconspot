package grpcapi

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"go.uber.org/zap"

	"github.com/begmaroman/beaconspot/beaconchain"
	beaconspotproto "github.com/begmaroman/beaconspot/proto/beaconspot"
	"github.com/begmaroman/beaconspot/proto/health"
)

// To make sure Handler implements grpcapiproto.BeaconSpotServiceHandler interface.
var _ beaconspotproto.BeaconSpotServiceServer = &handler{}

// Options serves as the dependency injection container to create a new handler.
type Options struct {
	BeaconChainClient beaconchain.BeaconChain
	Log               *zap.Logger
}

// handler implements grpcapiproto.BeaconSpotServiceHandler interface
type handler struct {
	beaconspotproto.UnimplementedBeaconSpotServiceServer

	beaconChainClient beaconchain.BeaconChain
	log               *zap.Logger
}

// New is the constructor of handler
func New(opts Options) beaconspotproto.BeaconSpotServiceServer {
	return &handler{
		beaconChainClient: opts.BeaconChainClient,
		log:               opts.Log,
	}
}

func (h *handler) SubnetsSubscribe(ctx context.Context, req *beaconspotproto.SubnetsSubscribeRequest) (*beaconspotproto.SubnetsSubscribeResponse, error) {
	subscription := make([]beaconchain.SubnetSubscription, len(req.GetSubscriptions()))
	for i, sub := range req.GetSubscriptions() {
		subscription[i] = beaconchain.SubnetSubscription{
			ValidatorIndex:   sub.GetValidatorIndex(),
			CommitteeIndex:   sub.GetCommitteeIndex(),
			CommitteesAtSlot: sub.GetCommitteesAtSlot(),
			Slot:             sub.GetSlot(),
			IsAggregator:     sub.GetIsAggregator(),
		}
	}

	if err := h.beaconChainClient.SubnetsSubscribe(ctx, subscription); err != nil {
		return nil, err
	}

	return &beaconspotproto.SubnetsSubscribeResponse{
		Result: &beaconspotproto.SubnetsSubscribeResponse_Empty{
			Empty: &empty.Empty{},
		},
	}, nil
}

func (h *handler) Health(context.Context, *empty.Empty) (*health.HealthResponse, error) {
	// TODO: Implement healthcheck
	return &health.HealthResponse{}, nil
}

func (h *handler) Ping(context.Context, *empty.Empty) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}
