package combined

import (
	"context"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/event"
	"github.com/pkg/errors"
	ethpb "github.com/prysmaticlabs/ethereumapis/eth/v1alpha1"

	"github.com/begmaroman/beaconspot/beaconchain"
)

// combined represents combined BeaconChain
type combined struct {
	beaconChains []beaconchain.BeaconChain
}

// New is the constructor of combined
func New(beaconChains ...beaconchain.BeaconChain) beaconchain.BeaconChain {
	return &combined{
		beaconChains: beaconChains,
	}
}

// GetAttestationData returns attestation data
func (c *combined) GetAttestationData(ctx context.Context, slot, index uint64) (*ethpb.AttestationData, error) {
	type result struct {
		attData *ethpb.AttestationData
		err     error
	}

	attestationDataFeed := &event.Feed{}

	attestationDataChan := make(chan *result)
	attestationDataSub := attestationDataFeed.Subscribe(attestationDataChan)
	defer attestationDataSub.Unsubscribe()

	for _, beaconChain := range c.beaconChains {
		go func(beaconChain beaconchain.BeaconChain) {
			attData, err := beaconChain.GetAttestationData(ctx, slot, index)
			attestationDataFeed.Send(&result{
				attData: attData,
				err:     err,
			})
		}(beaconChain)
	}

	var errs []string
	for i := 0; i < len(c.beaconChains); i++ {
		attData := <-attestationDataChan
		if attData.err != nil {
			errs = append(errs, attData.err.Error())
		} else {
			return attData.attData, nil
		}
	}

	return nil, errors.Errorf("failed to get attestation data from all nodes: %s", strings.Join(errs, ", "))
}

// ProposeAttestation proposes the given attestation
func (c *combined) ProposeAttestation(ctx context.Context, data *ethpb.AttestationData, aggregationBits, signature []byte) error {
	errs := make([]error, len(c.beaconChains))
	var wg sync.WaitGroup
	for i, beaconChain := range c.beaconChains {
		wg.Add(1)
		go func(i int, beaconChain beaconchain.BeaconChain) {
			defer wg.Done()
			errs[i] = beaconChain.ProposeAttestation(ctx, data, aggregationBits, signature)
		}(i, beaconChain)
	}
	wg.Wait()

	var errMsgs []string
	for _, err := range errs {
		if err == nil {
			return nil
		} else {
			errMsgs = append(errMsgs, err.Error())
		}
	}

	return errors.Errorf("failed to propose attestation data to all nodes: %s", strings.Join(errMsgs, ", "))
}

// GetBlock returns block by the given data
func (c *combined) GetBlock(ctx context.Context, slot uint64, randaoReveal, graffiti []byte) (*ethpb.BeaconBlock, error) {
	type result struct {
		block *ethpb.BeaconBlock
		err   error
	}

	dataFeed := &event.Feed{}

	dataChan := make(chan *result)
	dataSub := dataFeed.Subscribe(dataChan)
	defer dataSub.Unsubscribe()

	for _, beaconChain := range c.beaconChains {
		go func(beaconChain beaconchain.BeaconChain) {
			block, err := beaconChain.GetBlock(ctx, slot, randaoReveal, graffiti)
			dataFeed.Send(&result{
				block: block,
				err:   err,
			})
		}(beaconChain)
	}

	var errs []string
	for i := 0; i < len(c.beaconChains); i++ {
		data := <-dataChan
		if data.err != nil {
			errs = append(errs, data.err.Error())
		} else {
			return data.block, nil
		}
	}

	return nil, errors.Errorf("failed to get block from all nodes: %s", strings.Join(errs, ", "))
}

// ProposeBlock submits proposal for the given block
func (c *combined) ProposeBlock(ctx context.Context, signature []byte, block *ethpb.BeaconBlock) error {
	errs := make([]error, len(c.beaconChains))
	var wg sync.WaitGroup
	for i, beaconChain := range c.beaconChains {
		wg.Add(1)
		go func(i int, beaconChain beaconchain.BeaconChain) {
			defer wg.Done()
			errs[i] = beaconChain.ProposeBlock(ctx, signature, block)
		}(i, beaconChain)
	}
	wg.Wait()

	var errMsgs []string
	for _, err := range errs {
		if err == nil {
			return nil
		} else {
			errMsgs = append(errMsgs, err.Error())
		}
	}

	return errors.Errorf("failed to propose block to all nodes: %s", strings.Join(errMsgs, ", "))
}

// GetAggregateSelectionProof returns aggregated attestation
func (c *combined) GetAggregateSelectionProof(ctx context.Context, slot, committeeIndex uint64, publicKey, sig []byte) (*ethpb.AggregateAttestationAndProof, error) {
	type result struct {
		res *ethpb.AggregateAttestationAndProof
		err error
	}

	feed := &event.Feed{}

	dataChan := make(chan *result)
	dataSub := feed.Subscribe(dataChan)
	defer dataSub.Unsubscribe()

	for _, beaconChain := range c.beaconChains {
		go func(beaconChain beaconchain.BeaconChain) {
			res, err := beaconChain.GetAggregateSelectionProof(ctx, slot, committeeIndex, publicKey, sig)
			feed.Send(&result{
				res: res,
				err: err,
			})
		}(beaconChain)
	}

	var errs []string
	for i := 0; i < len(c.beaconChains); i++ {
		res := <-dataChan
		if res.err != nil {
			errs = append(errs, res.err.Error())
		} else {
			return res.res, nil
		}
	}

	return nil, errors.Errorf("failed to request aggregated attestation to all nodes: %s", strings.Join(errs, ", "))
}

// SubmitSignedAggregateSelectionProof verifies given aggregate and proofs and publishes them on appropriate gossipsub topic
func (c *combined) SubmitSignedAggregateSelectionProof(ctx context.Context, signature []byte, message *ethpb.AggregateAttestationAndProof) error {
	errs := make([]error, len(c.beaconChains))
	var wg sync.WaitGroup
	for i, beaconChain := range c.beaconChains {
		wg.Add(1)
		go func(i int, beaconChain beaconchain.BeaconChain) {
			defer wg.Done()
			errs[i] = beaconChain.SubmitSignedAggregateSelectionProof(ctx, signature, message)
		}(i, beaconChain)
	}
	wg.Wait()

	var errMsgs []string
	for _, err := range errs {
		if err == nil {
			return nil
		} else {
			errMsgs = append(errMsgs, err.Error())
		}
	}

	return errors.Errorf("failed to submit signed aggregation to all nodes: %s", strings.Join(errMsgs, ", "))
}

// SubnetsSubscribe subscribes on the given subnets
func (c *combined) SubnetsSubscribe(ctx context.Context, subscriptions []beaconchain.SubnetSubscription) error {
	errs := make([]error, len(c.beaconChains))
	var wg sync.WaitGroup
	for i, beaconChain := range c.beaconChains {
		wg.Add(1)
		go func(i int, beaconChain beaconchain.BeaconChain) {
			defer wg.Done()
			errs[i] = beaconChain.SubnetsSubscribe(ctx, subscriptions)
		}(i, beaconChain)
	}
	wg.Wait()

	var errMsgs []string
	for _, err := range errs {
		if err == nil {
			return nil
		} else {
			errMsgs = append(errMsgs, err.Error())
		}
	}

	return errors.Errorf("failed to subscribe on subnets to all nodes: %s", strings.Join(errMsgs, ", "))
}

// DomainData returns domain data by the given request
func (c *combined) DomainData(ctx context.Context, epoch uint64, domain []byte) ([]byte, error) {
	type result struct {
		res []byte
		err error
	}

	feed := &event.Feed{}

	dataChan := make(chan *result)
	dataSub := feed.Subscribe(dataChan)
	defer dataSub.Unsubscribe()

	for _, beaconChain := range c.beaconChains {
		go func(beaconChain beaconchain.BeaconChain) {
			res, err := beaconChain.DomainData(ctx, epoch, domain)
			feed.Send(&result{
				res: res,
				err: err,
			})
		}(beaconChain)
	}

	var errs []string
	for i := 0; i < len(c.beaconChains); i++ {
		res := <-dataChan
		if res.err != nil {
			errs = append(errs, res.err.Error())
		} else {
			return res.res, nil
		}
	}

	return nil, errors.Errorf("failed to get domain data from all nodes: %s", strings.Join(errs, ", "))
}

// StreamDuties returns client to stream duties
func (c *combined) StreamDuties(ctx context.Context, pubKeys [][]byte) (ethpb.BeaconNodeValidator_StreamDutiesClient, error) {
	var wg sync.WaitGroup
	clients := make([]ethpb.BeaconNodeValidator_StreamDutiesClient, len(c.beaconChains))
	errs := make([]error, len(c.beaconChains))
	for i, beaconChain := range c.beaconChains {
		wg.Add(1)
		go func(i int, beaconChain beaconchain.BeaconChain) {
			clients[i], errs[i] = beaconChain.StreamDuties(ctx, pubKeys)
			wg.Done()
		}(i, beaconChain)
	}
	wg.Wait()

	for _, err := range errs {
		if err != nil {
			return nil, errors.Wrap(err, "failed to stream duties")
		}
	}

	return newStreamDuties(clients...), nil
}

// GetGenesis returns genesis data
func (c *combined) GetGenesis(ctx context.Context) (*ethpb.Genesis, error) {
	type result struct {
		res *ethpb.Genesis
		err error
	}

	feed := &event.Feed{}

	dataChan := make(chan *result)
	dataSub := feed.Subscribe(dataChan)
	defer dataSub.Unsubscribe()

	for _, beaconChain := range c.beaconChains {
		go func(beaconChain beaconchain.BeaconChain) {
			res, err := beaconChain.GetGenesis(ctx)
			feed.Send(&result{
				res: res,
				err: err,
			})
		}(beaconChain)
	}

	var errs []string
	for i := 0; i < len(c.beaconChains); i++ {
		res := <-dataChan
		if res.err != nil {
			errs = append(errs, res.err.Error())
		} else {
			return res.res, nil
		}
	}

	return nil, errors.Errorf("failed to get genesis data from all nodes: %s", strings.Join(errs, ", "))
}
