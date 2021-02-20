package grpcapi

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"

	grpcapiproto "github.com/begmaroman/beaconspot/proto/beaconspot"
)

func (h *handler) GetAttestation(ctx context.Context, req *grpcapiproto.GetAttestationRequest, resp *grpcapiproto.GetAttestationResponse) error {
	attestationData, err := h.beaconChainClient.GetAttestationData(ctx, req.GetSlot(), req.GetCommitteeIndex())
	if err != nil {
		return err
	}

	resp.Result = &grpcapiproto.GetAttestationResponse_AttestationData{
		AttestationData: attestationData,
	}
	return nil
}

func (h *handler) ProposeAttestation(ctx context.Context, req *grpcapiproto.ProposeAttestationRequest, resp *grpcapiproto.ProposeAttestationResponse) error {
	err := h.beaconChainClient.ProposeAttestation(ctx, req.GetAttestationData(), req.GetAggregationBits(), req.GetSignature())
	if err != nil {
		return err
	}

	resp.Result = &grpcapiproto.ProposeAttestationResponse_Empty{
		Empty: &empty.Empty{},
	}
	return nil
}
