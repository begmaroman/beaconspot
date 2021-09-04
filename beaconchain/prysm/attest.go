package prysm

import (
	"context"

	"github.com/pkg/errors"
	types "github.com/prysmaticlabs/eth2-types"
	ethpb "github.com/prysmaticlabs/prysm/proto/prysm/v1alpha1"
)

// GetAttestationData returns attestation data
func (c *prysmGRPC) GetAttestationData(ctx context.Context, slot types.Slot, committeeIndex types.CommitteeIndex) (*ethpb.AttestationData, error) {
	resp, err := c.validatorClient.GetAttestationData(ctx, &ethpb.AttestationDataRequest{
		Slot:           slot,
		CommitteeIndex: committeeIndex,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "Prysm: failed to get attestation data")
	}

	c.logger.Debug("got attestation data from Prysm")

	return resp, nil
}

// ProposeAttestation proposes the given attestation
func (c *prysmGRPC) ProposeAttestation(ctx context.Context, data *ethpb.AttestationData, aggregationBits, signature []byte) error {
	_, err := c.validatorClient.ProposeAttestation(ctx, &ethpb.Attestation{
		AggregationBits: aggregationBits,
		Data:            data,
		Signature:       signature,
	})
	if err != nil {
		return errors.Wrap(err, "Prysm: failed to propose attestation")
	}

	return nil
}
