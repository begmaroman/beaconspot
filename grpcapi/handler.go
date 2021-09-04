package grpcapi

import (
	"context"

	types "github.com/prysmaticlabs/eth2-types"

	"github.com/golang/protobuf/ptypes/empty"
	"go.uber.org/zap"

	"github.com/begmaroman/beaconspot/beaconchain"
	beaconspotproto "github.com/begmaroman/beaconspot/proto/beaconspot"
	"github.com/begmaroman/beaconspot/proto/health"
	proto "github.com/begmaroman/beaconspot/proto/status"
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

func (h *handler) DomainData(ctx context.Context, req *beaconspotproto.DomainDataRequest) (*beaconspotproto.DomainDataResponse, error) {
	domainData, err := h.beaconChainClient.DomainData(ctx, types.Epoch(req.GetEpoch()), req.GetDomain())
	if err != nil {
		return nil, err
	}

	return &beaconspotproto.DomainDataResponse{
		Result: &beaconspotproto.DomainDataResponse_DomainData{
			DomainData: domainData,
		},
	}, nil
}

func (h *handler) SubnetsSubscribe(ctx context.Context, req *beaconspotproto.SubnetsSubscribeRequest) (*beaconspotproto.SubnetsSubscribeResponse, error) {
	subscription := make([]beaconchain.SubnetSubscription, len(req.GetSubscriptions()))
	for i, sub := range req.GetSubscriptions() {
		subscription[i] = beaconchain.SubnetSubscription{
			ValidatorIndex:   types.ValidatorIndex(sub.GetValidatorIndex()),
			CommitteeIndex:   types.CommitteeIndex(sub.GetCommitteeIndex()),
			CommitteesAtSlot: sub.GetCommitteesAtSlot(),
			Slot:             types.Slot(sub.GetSlot()),
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

func (h *handler) StreamDuties(req *beaconspotproto.StreamDutiesRequest, srv beaconspotproto.BeaconSpotService_StreamDutiesServer) error {
	client, err := h.beaconChainClient.StreamDuties(srv.Context(), req.GetPublicKeys())
	if err != nil {
		return err
	}

	for {
		select {
		case <-srv.Context().Done():
			client.CloseSend()
			return nil
		default:
			resp, err := client.Recv()
			if err != nil {
				if err := srv.Send(&beaconspotproto.StreamDutiesResponse{
					Result: &beaconspotproto.StreamDutiesResponse_Error{
						Error: &proto.Status{
							Message: err.Error(),
							Code:    500,
							Details: []byte("failed to receive duties"),
						},
					},
				}); err != nil {
					h.log.Error("failed to send error message", zap.Error(err))
				}
				break
			}

			if err := srv.Send(&beaconspotproto.StreamDutiesResponse{
				Result: &beaconspotproto.StreamDutiesResponse_Duties{
					Duties: resp,
				},
			}); err != nil {
				h.log.Error("failed to send duties message", zap.Error(err))
			}
		}
	}
}

func (h *handler) Health(context.Context, *empty.Empty) (*health.HealthResponse, error) {
	// TODO: Implement healthcheck
	return &health.HealthResponse{}, nil
}

func (h *handler) Ping(context.Context, *empty.Empty) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}
