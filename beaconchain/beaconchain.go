package beaconchain

import (
	"context"

	types "github.com/prysmaticlabs/eth2-types"
	ethpb "github.com/prysmaticlabs/prysm/proto/prysm/v1alpha1"
)

// SubnetSubscription contains data to subscribe on subnets
type SubnetSubscription struct {
	ValidatorIndex   types.ValidatorIndex
	CommitteeIndex   types.CommitteeIndex
	CommitteesAtSlot uint64
	Slot             types.Slot
	IsAggregator     bool
}

// BeaconChain represents beacon chain behavior
type BeaconChain interface {
	// GetAttestationData returns attestation data by the given data
	GetAttestationData(ctx context.Context, slot types.Slot, committeeIndex types.CommitteeIndex) (*ethpb.AttestationData, error)

	// ProposeAttestation proposed the given attestation
	ProposeAttestation(ctx context.Context, data *ethpb.AttestationData, aggregationBits, signature []byte) error

	// GetBlock returns block by the given request
	GetBlock(ctx context.Context, slot types.Slot, randaoReveal, graffiti []byte) (*ethpb.BeaconBlock, error)

	// ProposeBlock proposes block
	ProposeBlock(ctx context.Context, signature []byte, block *ethpb.BeaconBlock) error

	// GetAggregateSelectionProof returns aggregated attestation
	GetAggregateSelectionProof(ctx context.Context, slot types.Slot, committeeIndex types.CommitteeIndex, publicKey, sig []byte) (*ethpb.AggregateAttestationAndProof, error)

	// SubmitSignedAggregateSelectionProof verifies given aggregate and proofs and publishes them on appropriate gossipsub topic
	SubmitSignedAggregateSelectionProof(ctx context.Context, signature []byte, message *ethpb.AggregateAttestationAndProof) error

	// SubnetsSubscribe subscribes on the given subnets
	SubnetsSubscribe(ctx context.Context, subscriptions []SubnetSubscription) error

	// DomainData returns domain data by the given request
	DomainData(ctx context.Context, epoch types.Epoch, domain []byte) ([]byte, error)

	// StreamDuties returns client to stream duties
	StreamDuties(ctx context.Context, pubKeys [][]byte) (ethpb.BeaconNodeValidator_StreamDutiesClient, error)

	// GetGenesis returns genesis data
	GetGenesis(ctx context.Context) (*ethpb.Genesis, error)
}
