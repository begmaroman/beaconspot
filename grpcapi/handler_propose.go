package grpcapi

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	types "github.com/prysmaticlabs/eth2-types"

	beaconspotproto "github.com/begmaroman/beaconspot/proto/beaconspot"
)

func (h *handler) GetBlock(ctx context.Context, req *beaconspotproto.GetBlockRequest) (*beaconspotproto.GetBlockResponse, error) {
	block, err := h.beaconChainClient.GetBlock(ctx, types.Slot(req.GetSlot()), req.GetRandaoReveal(), req.GetGraffiti())
	if err != nil {
		return nil, err
	}

	return &beaconspotproto.GetBlockResponse{
		Result: &beaconspotproto.GetBlockResponse_BeaconBlock{
			BeaconBlock: block,
		},
	}, nil
}

func (h *handler) ProposeBlock(ctx context.Context, req *beaconspotproto.ProposeBlockRequest) (*beaconspotproto.ProposeBlockResponse, error) {
	if err := h.beaconChainClient.ProposeBlock(ctx, req.GetSignature(), req.GetBeaconBlock()); err != nil {
		return nil, err
	}

	return &beaconspotproto.ProposeBlockResponse{
		Result: &beaconspotproto.ProposeBlockResponse_Empty{
			Empty: &empty.Empty{},
		},
	}, nil
}
