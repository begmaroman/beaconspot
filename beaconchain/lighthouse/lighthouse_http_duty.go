package lighthouse

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/bloxapp/eth2-key-manager/core"
	ethpb "github.com/prysmaticlabs/ethereumapis/eth/v1alpha1"
	"github.com/prysmaticlabs/prysm/endtoend/helpers"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"

	"github.com/begmaroman/beaconspot/utils/httpex"
	"github.com/begmaroman/beaconspot/utils/slicex"
)

// lighthouseStreamDuties implements ethpb.BeaconNodeValidator_StreamDutiesClient interface using LightHouse HTTP API
type lighthouseStreamDuties struct {
	ctx              context.Context
	addr             string
	client           *http.Client
	validatorIndexes []string
	logger           *zap.Logger
	epochTicker      *helpers.EpochTicker

	errCh    chan error
	dutiesCh chan *ethpb.DutiesResponse
}

// newLighthouseStreamDuties is the constructor of lighthouseStreamDuties
func newLighthouseStreamDuties(ctx context.Context, logger *zap.Logger, network core.Network, addr string, validatorIndexes []string) ethpb.BeaconNodeValidator_StreamDutiesClient {
	capacity := len(validatorIndexes) * 10 // TODO: Figure out how many proposal slots can be pre public key
	genesisTime := time.Unix(int64(network.MinGenesisTime()), 0)
	dutiesCh := make(chan *ethpb.DutiesResponse, capacity)
	errCh := make(chan error, capacity)
	epochTicker := helpers.GetEpochTicker(genesisTime, network.SlotsPerEpoch()*uint64(network.SlotDurationSec().Seconds()))

	client := &lighthouseStreamDuties{
		ctx:              ctx,
		addr:             addr,
		client:           httpex.CreateClient(),
		validatorIndexes: validatorIndexes,
		logger:           logger,
		epochTicker:      epochTicker,
		dutiesCh:         dutiesCh,
		errCh:            errCh,
	}

	go func() {
		for epoch := range epochTicker.C() {
			logger := logger.With(zap.Uint64("current_epoch", epoch), zap.Uint64("next_epoch", epoch+1))

			var currentEpochDuties []*ethpb.DutiesResponse_Duty
			var nextEpochDuties []*ethpb.DutiesResponse_Duty
			errs := make([]error, 2)

			var wg sync.WaitGroup

			// Get current epoch duties
			logger.Info("loading duties for the current epoch...")
			wg.Add(1)
			go func() {
				defer wg.Done()

				duties, err := client.getDutiesForEpoch(epoch)
				if err != nil {
					errs[0] = err
					return
				}

				currentEpochDuties = duties
				logger.Info("loaded duties for the current epoch")
			}()

			// Get next epoch duties
			logger.Info("loading duties for the next epoch...")
			wg.Add(1)
			go func() {
				defer wg.Done()

				duties, err := client.getDutiesForEpoch(epoch + 1)
				if err != nil {
					errs[1] = err
					return
				}

				nextEpochDuties = duties
				logger.Info("loaded duties for the next epoch")
			}()

			wg.Wait()

			var foundErr bool
			for _, err := range errs {
				if err != nil {
					errCh <- err
					foundErr = true
					break
				}
			}

			if foundErr {
				continue
			}

			logger.Info("sending duties result to the channel...")
			dutiesCh <- &ethpb.DutiesResponse{
				CurrentEpochDuties: currentEpochDuties,
				NextEpochDuties:    nextEpochDuties,
			}
			logger.Info("duties result sent to the channel")
		}
	}()

	return client
}

func (c *lighthouseStreamDuties) Recv() (*ethpb.DutiesResponse, error) {
	return nil, errors.New("not implemented yet")

	select {
	case duty := <-c.dutiesCh:
		return duty, nil
	case err := <-c.errCh:
		return nil, err
	}
}

func (c *lighthouseStreamDuties) Header() (metadata.MD, error) {
	panic("not implemented")
}

func (c *lighthouseStreamDuties) Trailer() metadata.MD {
	panic("not implemented")
}

func (c *lighthouseStreamDuties) CloseSend() error {
	c.epochTicker.Done()
	return nil
}

func (c *lighthouseStreamDuties) Context() context.Context {
	return c.ctx
}

func (c *lighthouseStreamDuties) SendMsg(m interface{}) error {
	panic("not implemented")
}

func (c *lighthouseStreamDuties) RecvMsg(m interface{}) error {
	panic("not implemented")
}

func (c *lighthouseStreamDuties) getDutiesForEpoch(epoch uint64) ([]*ethpb.DutiesResponse_Duty, error) {
	logger := c.logger.With(zap.Uint64("epoch", epoch))

	var attesterDuties []*ethpb.DutiesResponse_Duty
	var proposerDuties []*ethpb.DutiesResponse_Duty

	var wg sync.WaitGroup
	errs := make([]error, 2)

	// Get attester duties
	logger.Info("loading attester duties...")
	wg.Add(1)
	go func() {
		defer wg.Done()

		url := fmt.Sprintf("%s/eth/v1/validator/duties/attester/%d", c.addr, epoch)
		reqBody := fmt.Sprintf(`["%s"]`, strings.Join(c.validatorIndexes, `","`))

		req, err := http.NewRequestWithContext(c.ctx, http.MethodPost, url, strings.NewReader(reqBody))
		if err != nil {
			errs[0] = err
			return
		}

		resp, err := c.client.Do(req)
		if err != nil {
			errs[0] = err
			return
		}

		defer resp.Body.Close()

		var duties dutyModel
		if err := json.NewDecoder(resp.Body).Decode(&duties); err != nil {
			errs[0] = err
			return
		}

		for _, duty := range duties.Data {
			if !slicex.ContainsString(c.validatorIndexes, duty.ValidatorIndex) {
				continue
			}

			protoModel, err := duty.toAttesterDuties()
			if err != nil {
				errs[1] = err
				return
			}

			attesterDuties = append(attesterDuties, protoModel)
		}

		logger.Info("loaded attester duties")
	}()

	// Get proposer duties
	logger.Info("loading proposer duties...")
	wg.Add(1)
	go func() {
		defer wg.Done()

		url := fmt.Sprintf("%s/eth/v1/validator/duties/proposer/%d", c.addr, epoch)

		req, err := http.NewRequestWithContext(c.ctx, http.MethodGet, url, nil)
		if err != nil {
			errs[1] = err
			return
		}

		resp, err := c.client.Do(req)
		if err != nil {
			errs[1] = err
			return
		}

		defer resp.Body.Close()

		var duties dutyModel
		if err := json.NewDecoder(resp.Body).Decode(&duties); err != nil {
			errs[1] = err
			return
		}

		for _, duty := range duties.Data {
			if !slicex.ContainsString(c.validatorIndexes, duty.ValidatorIndex) {
				continue
			}

			protoModel, err := duty.toProposerDuties()
			if err != nil {
				errs[1] = err
				return
			}

			proposerDuties = append(proposerDuties, protoModel)
		}

		logger.Info("loaded proposer duties")
	}()

	wg.Wait()

	var foundErr bool
	for _, err := range errs {
		if err != nil {
			return nil, err
		}
	}

	if foundErr {
		return nil, nil
	}

	return append(attesterDuties, proposerDuties...), nil
}
