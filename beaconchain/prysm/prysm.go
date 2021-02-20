package prysm

import (
	"context"
	"encoding/hex"

	types "github.com/prysmaticlabs/eth2-types"

	"github.com/pkg/errors"
	ethpb "github.com/prysmaticlabs/ethereumapis/eth/v1alpha1"
	"go.uber.org/zap"

	"github.com/begmaroman/beaconspot/beaconchain"
)

// prysmGRPC implements beaconchain.BeaconChain interface using Prysm beacon node via gRPC
type prysmGRPC struct {
	validatorClient ethpb.BeaconNodeValidatorClient
	logger          *zap.Logger
}

// New is the constructor of prysmGRPC
func New(logger *zap.Logger, validatorClient ethpb.BeaconNodeValidatorClient) beaconchain.BeaconChain {
	return &prysmGRPC{
		validatorClient: validatorClient,
		logger:          logger,
	}
}

// GetAttestationData returns attestation data
func (c *prysmGRPC) GetAttestationData(ctx context.Context, slot, committeeIndex uint64) (*ethpb.AttestationData, error) {
	resp, err := c.validatorClient.GetAttestationData(ctx, &ethpb.AttestationDataRequest{
		Slot:           types.Slot(slot),
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

// GetBlock returns block by the given data
func (c *prysmGRPC) GetBlock(ctx context.Context, slot uint64, randaoReveal, graffiti []byte) (*ethpb.BeaconBlock, error) {
	b, err := c.validatorClient.GetBlock(ctx, &ethpb.BlockRequest{
		Slot:         types.Slot(slot),
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

// SubmitAggregateSelectionProof returns aggregated attestation
func (c *prysmGRPC) SubmitAggregateSelectionProof(ctx context.Context, slot, committeeIndex uint64, publicKey, sig []byte) (*ethpb.AggregateAttestationAndProof, error) {
	res, err := c.validatorClient.SubmitAggregateSelectionProof(ctx, &ethpb.AggregateSelectionRequest{
		Slot:           types.Slot(slot),
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

// SubnetsSubscribe subscribes on the given subnets
func (c *prysmGRPC) SubnetsSubscribe(ctx context.Context, subscriptions []beaconchain.SubnetSubscription) error {
	subscribeReq := &ethpb.CommitteeSubnetsSubscribeRequest{
		Slots:        make([]types.Slot, len(subscriptions)),
		CommitteeIds: make([]uint64, len(subscriptions)),
		IsAggregator: make([]bool, len(subscriptions)),
	}

	for i, subscription := range subscriptions {
		subscribeReq.Slots[i] = types.Slot(subscription.Slot)
		subscribeReq.CommitteeIds[i] = subscription.CommitteeIndex
		subscribeReq.IsAggregator[i] = subscription.IsAggregator
	}

	_, err := c.validatorClient.SubscribeCommitteeSubnets(ctx, subscribeReq)
	if err != nil {
		return errors.Wrap(err, "Prysm: failed to subscribe on subnets")
	}

	return nil
}

// DomainData returns domain data by the given request
func (c *prysmGRPC) DomainData(ctx context.Context, epoch uint64, domain []byte) ([]byte, error) {
	res, err := c.validatorClient.DomainData(ctx, &ethpb.DomainRequest{
		Epoch:  types.Epoch(epoch),
		Domain: domain,
	})
	if err != nil {
		return nil, errors.Wrap(err, "Prysm: failed to get domain data")
	}

	c.logger.Info("Prysm: got domain data", zap.String("domain_data", hex.EncodeToString(res.GetSignatureDomain())))

	return res.GetSignatureDomain(), nil
}

// StreamDuties returns client to stream duties
func (c *prysmGRPC) StreamDuties(ctx context.Context, pubKeys [][]byte) (ethpb.BeaconNodeValidator_StreamDutiesClient, error) {
	res, err := c.validatorClient.StreamDuties(ctx, &ethpb.DutiesRequest{
		PublicKeys: pubKeys,
	})
	if err != nil {
		return nil, errors.Wrap(err, "Prysm: failed to get stream duties client")
	}

	return res, nil
}
