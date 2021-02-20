package beaconchain

import (
	"context"

	ethpb "github.com/prysmaticlabs/ethereumapis/eth/v1alpha1"
)

// SubnetSubscription contains data to subscribe on subnets
type SubnetSubscription struct {
	ValidatorIndex   uint64
	CommitteeIndex   uint64
	CommitteesAtSlot uint64
	Slot             uint64
	IsAggregator     bool
}

// BeaconChain represents beacon chain behavior
type BeaconChain interface {
	// GetAttestationData returns attestation data by the given data
	GetAttestationData(ctx context.Context, slot, committeeIndex uint64) (*ethpb.AttestationData, error)

	// ProposeAttestation proposed the given attestation
	ProposeAttestation(ctx context.Context, data *ethpb.AttestationData, aggregationBits, signature []byte) error

	// GetBlock returns block by the given request
	GetBlock(ctx context.Context, slot uint64, randaoReveal, graffiti []byte) (*ethpb.BeaconBlock, error)

	// ProposeBlock proposes block
	ProposeBlock(ctx context.Context, signature []byte, block *ethpb.BeaconBlock) error

	// SubmitAggregateSelectionProof returns aggregated attestation
	SubmitAggregateSelectionProof(ctx context.Context, slot, committeeIndex uint64, publicKey, sig []byte) (*ethpb.AggregateAttestationAndProof, error)

	// SubmitSignedAggregateSelectionProof verifies given aggregate and proofs and publishes them on appropriate gossipsub topic
	SubmitSignedAggregateSelectionProof(ctx context.Context, signature []byte, message *ethpb.AggregateAttestationAndProof) error

	// SubnetsSubscribe subscribes on the given subnets
	SubnetsSubscribe(ctx context.Context, subscriptions []SubnetSubscription) error

	// DomainData returns domain data by the given request
	DomainData(ctx context.Context, epoch uint64, domain []byte) ([]byte, error)

	// StreamDuties returns client to stream duties
	StreamDuties(ctx context.Context, pubKeys [][]byte) (ethpb.BeaconNodeValidator_StreamDutiesClient, error)
}
