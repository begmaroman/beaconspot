package prysm

import (
	"context"

	"github.com/pkg/errors"
	types "github.com/prysmaticlabs/eth2-types"
	ethpb "github.com/prysmaticlabs/prysm/proto/prysm/v1alpha1"
)

// GetBlock returns block by the given data
func (c *prysmGRPC) GetBlock(ctx context.Context, slot types.Slot, randaoReveal, graffiti []byte) (*ethpb.BeaconBlock, error) {
	b, err := c.validatorClient.GetBlock(ctx, &ethpb.BlockRequest{
		Slot:         slot,
		RandaoReveal: randaoReveal,
		Graffiti:     graffiti,
	})
	if err != nil {
		return nil, errors.Wrap(err, "Prysm: failed to get block")
	}

	return b, nil
}

// ProposeBlock submits proposal for the given block
func (c *prysmGRPC) ProposeBlock(ctx context.Context, signature []byte, block *ethpb.BeaconBlock) error {
	_, err := c.validatorClient.ProposeBlock(ctx, &ethpb.SignedBeaconBlock{
		Block:     block,
		Signature: signature,
	})
	if err != nil {
		return errors.Wrap(err, "Prysm: failed to propose block")
	}

	return nil
}
