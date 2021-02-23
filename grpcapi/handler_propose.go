package grpcapi

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"

	grpcapiproto "github.com/begmaroman/beaconspot/proto/beaconspot"
)

func (h *handler) GetBlock(ctx context.Context, req *grpcapiproto.GetBlockRequest) (*grpcapiproto.GetBlockResponse, error) {
	block, err := h.beaconChainClient.GetBlock(ctx, req.GetSlot(), req.GetRandaoReveal(), req.GetGraffiti())
	if err != nil {
		return nil, err
	}

	return &grpcapiproto.GetBlockResponse{
		Result: &grpcapiproto.GetBlockResponse_BeaconBlock{
			BeaconBlock: block,
		},
	}, nil
}

func (h *handler) ProposeBlock(ctx context.Context, req *grpcapiproto.ProposeBlockRequest) (*grpcapiproto.ProposeBlockResponse, error) {
	if err := h.beaconChainClient.ProposeBlock(ctx, req.GetSignature(), req.GetBeaconBlock()); err != nil {
		return nil, err
	}

	return &grpcapiproto.ProposeBlockResponse{
		Result: &grpcapiproto.ProposeBlockResponse_Empty{
			Empty: &empty.Empty{},
		},
	}, nil
}
