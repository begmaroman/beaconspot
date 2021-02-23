package grpcapi

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"

	beaconspotproto "github.com/begmaroman/beaconspot/proto/beaconspot"
)

func (h *handler) GetAggregateSelectionProof(ctx context.Context, req *beaconspotproto.GetAggregateSelectionProofRequest) (*beaconspotproto.GetAggregateSelectionProofResponse, error) {
	data, err := h.beaconChainClient.GetAggregateSelectionProof(ctx, req.GetSlot(), req.GetCommitteeIndex(), req.GetPublicKey(), req.GetSignature())
	if err != nil {
		return nil, err
	}

	return &beaconspotproto.GetAggregateSelectionProofResponse{
		Result: &beaconspotproto.GetAggregateSelectionProofResponse_Data{
			Data: data,
		},
	}, nil
}

func (h *handler) SubmitSignedAggregateSelectionProof(ctx context.Context, req *beaconspotproto.SubmitSignedAggregateSelectionProofRequest) (*beaconspotproto.SubmitSignedAggregateSelectionProofResponse, error) {
	if err := h.beaconChainClient.SubmitSignedAggregateSelectionProof(ctx, req.GetSignature(), req.GetData()); err != nil {
		return nil, err
	}

	return &beaconspotproto.SubmitSignedAggregateSelectionProofResponse{
		Result: &beaconspotproto.SubmitSignedAggregateSelectionProofResponse_Empty{
			Empty: &empty.Empty{},
		},
	}, nil
}
