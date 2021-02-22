package grpcapi

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"

	grpcapiproto "github.com/begmaroman/beaconspot/proto/beaconspot"
)

func (h *handler) GetAggregateSelectionProof(ctx context.Context, req *grpcapiproto.GetAggregateSelectionProofRequest) (*grpcapiproto.GetAggregateSelectionProofResponse, error) {
	data, err := h.beaconChainClient.GetAggregateSelectionProof(ctx, req.GetSlot(), req.GetCommitteeIndex(), req.GetPublicKey(), req.GetSignature())
	if err != nil {
		return nil, err
	}

	return &grpcapiproto.GetAggregateSelectionProofResponse{
		Result: &grpcapiproto.GetAggregateSelectionProofResponse_Data{
			Data: data,
		},
	}, nil
}

func (h *handler) SubmitSignedAggregateSelectionProof(ctx context.Context, req *grpcapiproto.SubmitSignedAggregateSelectionProofRequest) (*grpcapiproto.SubmitSignedAggregateSelectionProofResponse, error) {
	if err := h.beaconChainClient.SubmitSignedAggregateSelectionProof(ctx, req.GetSignature(), req.GetData()); err != nil {
		return nil, err
	}

	return &grpcapiproto.SubmitSignedAggregateSelectionProofResponse{
		Result: &grpcapiproto.SubmitSignedAggregateSelectionProofResponse_Empty{
			Empty: &empty.Empty{},
		},
	}, nil
}
