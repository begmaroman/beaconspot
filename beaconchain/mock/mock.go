package mock

import (
	"context"

	ethpb "github.com/prysmaticlabs/ethereumapis/eth/v1alpha1"

	"github.com/begmaroman/beaconspot/beaconchain"
)

// BeaconChain represents BeaconChain of BeaconChain
type BeaconChain struct {
	GetAttestationDataFn                  func(ctx context.Context, slot, index uint64) (*ethpb.AttestationData, error)
	ProposeAttestationFn                  func(ctx context.Context, data *ethpb.AttestationData, aggregationBits, signature []byte) error
	GetBlockFn                            func(ctx context.Context, slot uint64, randaoReveal, graffiti []byte) (*ethpb.BeaconBlock, error)
	ProposeBlockFn                        func(ctx context.Context, signature []byte, block *ethpb.BeaconBlock) error
	GetAggregateSelectionProofFn          func(ctx context.Context, slot, committeeIndex uint64, publicKey, sig []byte) (*ethpb.AggregateAttestationAndProof, error)
	SubmitSignedAggregateSelectionProofFn func(ctx context.Context, signature []byte, message *ethpb.AggregateAttestationAndProof) error
	SubnetsSubscribeFn                    func(ctx context.Context, subscriptions []beaconchain.SubnetSubscription) error
	DomainDataFn                          func(ctx context.Context, epoch uint64, domain []byte) ([]byte, error)
	StreamDutiesFn                        func(ctx context.Context, pubKeys [][]byte) (ethpb.BeaconNodeValidator_StreamDutiesClient, error)
	GetGenesisFn                          func(ctx context.Context) (*ethpb.Genesis, error)
}

// GetAttestationData returns attestation data
func (m *BeaconChain) GetAttestationData(ctx context.Context, slot, index uint64) (*ethpb.AttestationData, error) {
	return m.GetAttestationDataFn(ctx, slot, index)
}

// ProposeAttestation proposes the given attestation
func (m *BeaconChain) ProposeAttestation(ctx context.Context, data *ethpb.AttestationData, aggregationBits, signature []byte) error {
	return m.ProposeAttestationFn(ctx, data, aggregationBits, signature)
}

// GetBlock returns block by the given data
func (m *BeaconChain) GetBlock(ctx context.Context, slot uint64, randaoReveal, graffiti []byte) (*ethpb.BeaconBlock, error) {
	return m.GetBlockFn(ctx, slot, randaoReveal, graffiti)
}

// ProposeBlock submits proposal for the given block
func (m *BeaconChain) ProposeBlock(ctx context.Context, signature []byte, block *ethpb.BeaconBlock) error {
	return m.ProposeBlockFn(ctx, signature, block)
}

// GetAggregateSelectionProof returns aggregated attestation
func (m *BeaconChain) GetAggregateSelectionProof(ctx context.Context, slot, committeeIndex uint64, publicKey, sig []byte) (*ethpb.AggregateAttestationAndProof, error) {
	return m.GetAggregateSelectionProofFn(ctx, slot, committeeIndex, publicKey, sig)
}

// SubmitSignedAggregateSelectionProof verifies given aggregate and proofs and publishes them on appropriate gossipsub topic
func (m *BeaconChain) SubmitSignedAggregateSelectionProof(ctx context.Context, signature []byte, message *ethpb.AggregateAttestationAndProof) error {
	return m.SubmitSignedAggregateSelectionProofFn(ctx, signature, message)
}

// SubnetsSubscribe subscribes on the given subnets
func (m *BeaconChain) SubnetsSubscribe(ctx context.Context, subscriptions []beaconchain.SubnetSubscription) error {
	return m.SubnetsSubscribeFn(ctx, subscriptions)
}

// DomainData returns domain data by the given request
func (m *BeaconChain) DomainData(ctx context.Context, epoch uint64, domain []byte) ([]byte, error) {
	return m.DomainDataFn(ctx, epoch, domain)
}

// StreamDuties returns client to stream duties
func (m *BeaconChain) StreamDuties(ctx context.Context, pubKeys [][]byte) (ethpb.BeaconNodeValidator_StreamDutiesClient, error) {
	return m.StreamDutiesFn(ctx, pubKeys)
}

// GetGenesis returns genesis data
func (m *BeaconChain) GetGenesis(ctx context.Context) (*ethpb.Genesis, error) {
	return m.GetGenesisFn(ctx)
}
