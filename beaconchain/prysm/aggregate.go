package prysm

import (
	"context"

	"github.com/pkg/errors"
	ethpb "github.com/prysmaticlabs/ethereumapis/eth/v1alpha1"
)

// GetAggregateSelectionProof returns aggregated attestation
func (c *prysmGRPC) GetAggregateSelectionProof(ctx context.Context, slot, committeeIndex uint64, publicKey, sig []byte) (*ethpb.AggregateAttestationAndProof, error) {
	res, err := c.validatorClient.SubmitAggregateSelectionProof(ctx, &ethpb.AggregateSelectionRequest{
		Slot:           slot,
		CommitteeIndex: committeeIndex,
		PublicKey:      publicKey,
		SlotSignature:  sig,
	})
	if err != nil {
		return nil, errors.Wrap(err, "Prysm: failed to submit aggregation")
	}

	return res.GetAggregateAndProof(), nil
}

// SubmitSignedAggregateSelectionProof verifies given aggregate and proofs and publishes them on appropriate gossipsub topic
func (c *prysmGRPC) SubmitSignedAggregateSelectionProof(ctx context.Context, signature []byte, message *ethpb.AggregateAttestationAndProof) error {
	_, err := c.validatorClient.SubmitSignedAggregateSelectionProof(ctx, &ethpb.SignedAggregateSubmitRequest{
		SignedAggregateAndProof: &ethpb.SignedAggregateAttestationAndProof{
			Message:   message,
			Signature: signature,
		},
	})
	if err != nil {
		return errors.Wrap(err, "Prysm: failed to submit signed aggregation")
	}

	return nil
}
