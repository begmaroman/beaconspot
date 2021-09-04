package grpcapi

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	types "github.com/prysmaticlabs/eth2-types"

	beaconspotproto "github.com/begmaroman/beaconspot/proto/beaconspot"
)

func (h *handler) GetAttestation(ctx context.Context, req *beaconspotproto.GetAttestationRequest) (*beaconspotproto.GetAttestationResponse, error) {
	attestationData, err := h.beaconChainClient.GetAttestationData(ctx, types.Slot(req.GetSlot()), types.CommitteeIndex(req.GetCommitteeIndex()))
	if err != nil {
		return nil, err
	}

	return &beaconspotproto.GetAttestationResponse{
		Result: &beaconspotproto.GetAttestationResponse_AttestationData{
			AttestationData: attestationData,
		},
	}, nil
}

func (h *handler) ProposeAttestation(ctx context.Context, req *beaconspotproto.ProposeAttestationRequest) (*beaconspotproto.ProposeAttestationResponse, error) {
	err := h.beaconChainClient.ProposeAttestation(ctx, req.GetAttestationData(), req.GetAggregationBits(), req.GetSignature())
	if err != nil {
		return nil, err
	}

	return &beaconspotproto.ProposeAttestationResponse{
		Result: &beaconspotproto.ProposeAttestationResponse_Empty{
			Empty: &empty.Empty{},
		},
	}, nil
}
