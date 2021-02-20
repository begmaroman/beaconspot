package grpcapi

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"

	grpcapiproto "github.com/begmaroman/beaconspot/proto/beaconspot"
)

func (h *handler) GetAttestation(ctx context.Context, req *grpcapiproto.GetAttestationRequest) (*grpcapiproto.GetAttestationResponse, error) {
	attestationData, err := h.beaconChainClient.GetAttestationData(ctx, req.GetSlot(), req.GetCommitteeIndex())
	if err != nil {
		return nil, err
	}

	return &grpcapiproto.GetAttestationResponse{
		Result: &grpcapiproto.GetAttestationResponse_AttestationData{
			AttestationData: attestationData,
		},
	}, nil
}

func (h *handler) ProposeAttestation(ctx context.Context, req *grpcapiproto.ProposeAttestationRequest) (*grpcapiproto.ProposeAttestationResponse, error) {
	err := h.beaconChainClient.ProposeAttestation(ctx, req.GetAttestationData(), req.GetAggregationBits(), req.GetSignature())
	if err != nil {
		return nil, err
	}

	return &grpcapiproto.ProposeAttestationResponse{
		Result: &grpcapiproto.ProposeAttestationResponse_Empty{
			Empty: &empty.Empty{},
		},
	}, nil
}
