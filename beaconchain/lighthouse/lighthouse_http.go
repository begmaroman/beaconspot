package lighthouse

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

	"github.com/bloxapp/eth2-key-manager/core"
	"github.com/pkg/errors"
	ethpb "github.com/prysmaticlabs/ethereumapis/eth/v1alpha1"
	"github.com/prysmaticlabs/prysm/beacon-chain/core/helpers"
	pb "github.com/prysmaticlabs/prysm/proto/beacon/p2p/v1"
	"github.com/prysmaticlabs/prysm/shared/bytesutil"
	"go.uber.org/zap"

	"github.com/begmaroman/beaconspot/beaconchain"
	"github.com/begmaroman/beaconspot/utils/httpex"
)

// lighthouseHTTP implements BeaconChain interface using LightHouse beacon node via HTTP
type lighthouseHTTP struct {
	network    core.Network
	httpClient *http.Client
	addr       string
	logger     *zap.Logger

	genesisData     genesisModel
	genesisDataLock sync.Mutex
}

// New is the constructor of lighthouseHTTP
func New(logger *zap.Logger, network core.Network, addr string) beaconchain.BeaconChain {
	return &lighthouseHTTP{
		network:    network,
		httpClient: httpex.CreateClient(),
		addr:       addr,
		logger:     logger,
	}
}

// GetAttestationData returns attestation data
func (n *lighthouseHTTP) GetAttestationData(ctx context.Context, slot, index uint64) (*ethpb.AttestationData, error) {
	url := fmt.Sprintf("%s/eth/v1/validator/attestation_data?slot=%d&committee_index=%d", n.addr, slot, index)
	n.logger.Info("LightHouse: generated URL to get attestation data", zap.String("url", url))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "LightHouse: failed to create request with context")
	}

	resp, err := n.httpClient.Do(req)
	if err != nil {
		if err == context.DeadlineExceeded || err == context.Canceled {
			return nil, nil
		}

		n.logger.Error("LightHouse: failed to send request", zap.Error(err))
		return nil, errors.Wrap(err, "LightHouse: failed to send request")
	}

	if resp.StatusCode > 299 {
		n.logger.Error("LightHouse: unexpected response code", zap.Int("status_code", resp.StatusCode), zap.String("body", getResponseBodyRaw(resp)))
		return nil, errors.Errorf("LightHouse: unexpected response code %d", resp.StatusCode)
	}

	if resp.Body == nil {
		return nil, errors.New("LightHouse: empty response body")
	}

	defer resp.Body.Close()

	var attData attestationModel
	if err := json.NewDecoder(resp.Body).Decode(&attData); err != nil {
		n.logger.Error("LightHouse: failed to decode response body", zap.Error(err))
		return nil, errors.Wrap(err, "LightHouse: failed to decode response body")
	}

	n.logger.Debug("got attestation data from LightHouse")

	return attData.Data.toProto()
}

// ProposeAttestation proposes the given attestation
func (n *lighthouseHTTP) ProposeAttestation(_ context.Context, data *ethpb.AttestationData, aggregationBits, signature []byte) error {
	reqBody, err := json.Marshal([]attestationModel{toAttestationModel(&ethpb.Attestation{
		AggregationBits: aggregationBits,
		Data:            data,
		Signature:       signature,
	})})
	if err != nil {
		return errors.Wrap(err, "LightHouse: failed to marshal lighthouse JSON model")
	}

	req, err := http.NewRequest(http.MethodPost, n.addr+"/eth/v1/beacon/pool/attestations", bytes.NewReader(reqBody))
	if err != nil {
		return errors.Wrap(err, "LightHouse: failed to create request with context")
	}

	resp, err := n.httpClient.Do(req)
	if err != nil {
		n.logger.Error("LightHouse: failed to send request", zap.Error(err))
		return errors.Wrap(err, "LightHouse: failed to send request")
	}

	if resp.StatusCode > 299 {
		n.logger.Error("LightHouse: unexpected response code", zap.Int("status_code", resp.StatusCode), zap.String("body", getResponseBodyRaw(resp)))
		return errors.Errorf("LightHouse: unexpected response code %d", resp.StatusCode)
	}

	return nil
}

// GetBlock returns block by the given data
func (n *lighthouseHTTP) GetBlock(ctx context.Context, slot uint64, randaoReveal, graffiti []byte) (*ethpb.BeaconBlock, error) {
	url := fmt.Sprintf("%s/eth/v1/validator/blocks/%d?randao_reveal=%s&graffiti=%s",
		n.addr, slot, "0x"+hex.EncodeToString(randaoReveal), "0x"+hex.EncodeToString(graffiti))
	n.logger.Info("LightHouse: generated URL to get block", zap.String("url", url))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "LightHouse: failed to create request with context")
	}

	resp, err := n.httpClient.Do(req)
	if err != nil {
		if err == context.DeadlineExceeded || err == context.Canceled {
			return nil, nil
		}

		n.logger.Error("LightHouse: failed to send request", zap.Error(err))
		return nil, errors.Wrap(err, "LightHouse: failed to send request")
	}

	if resp.StatusCode > 299 {
		n.logger.Error("LightHouse: unexpected response code", zap.Int("status_code", resp.StatusCode), zap.String("body", getResponseBodyRaw(resp)))
		return nil, errors.Errorf("LightHouse: unexpected response code %d", resp.StatusCode)
	}

	if resp.Body == nil {
		return nil, errors.New("LightHouse: empty response body")
	}

	defer resp.Body.Close()

	var attData blockReqModel
	if err := json.NewDecoder(resp.Body).Decode(&attData); err != nil {
		n.logger.Error("LightHouse: failed to decode response body", zap.Error(err))
		return nil, errors.Wrap(err, "LightHouse: failed to decode response body")
	}

	n.logger.Debug("got block from LightHouse")

	return attData.Data.toProto()
}

// ProposeBlock submits proposal for the given block
func (n *lighthouseHTTP) ProposeBlock(_ context.Context, signature []byte, block *ethpb.BeaconBlock) error {
	reqBody, err := json.Marshal(toBlockSubmitModel(signature, block))
	if err != nil {
		return errors.Wrap(err, "LightHouse: failed to marshal lighthouse JSON model")
	}

	req, err := http.NewRequest(http.MethodPost, n.addr+"/eth/v1/beacon/blocks", bytes.NewReader(reqBody))
	if err != nil {
		return errors.Wrap(err, "LightHouse: failed to create request with context")
	}

	resp, err := n.httpClient.Do(req)
	if err != nil {
		n.logger.Error("LightHouse: failed to send request", zap.Error(err))
		return errors.Wrap(err, "LightHouse: failed to send request")
	}

	if resp.StatusCode > 299 {
		n.logger.Error("LightHouse: unexpected response code", zap.Int("status_code", resp.StatusCode), zap.String("body", getResponseBodyRaw(resp)))
		return errors.Errorf("LightHouse: unexpected response code %d", resp.StatusCode)
	}

	return nil
}

// GetAggregateSelectionProof returns aggregated attestation
func (n *lighthouseHTTP) GetAggregateSelectionProof(ctx context.Context, slot, committeeIndex uint64, publicKey, sig []byte) (*ethpb.AggregateAttestationAndProof, error) {
	attestationData, err := n.GetAttestationData(ctx, slot, committeeIndex)
	if err != nil {
		n.logger.Error("LightHouse: failed to get attestation data", zap.Error(err))
		return nil, errors.Wrap(err, "LightHouse: failed to get attestation data")
	}

	n.logger.Info("got attestation from LightHouse", zap.Any("att_data", attestationData))

	attestationDataRoot, err := attestationData.HashTreeRoot()
	if err != nil {
		n.logger.Error("LightHouse: failed to get attestation data root", zap.Error(err))
		return nil, errors.Wrap(err, "LightHouse: failed to get attestation data root")
	}

	url := fmt.Sprintf("%s/eth/v1/validator/aggregate_attestation?slot=%d&attestation_data_root=0x%s",
		n.addr, slot, hex.EncodeToString(attestationDataRoot[:]))
	n.logger.Info("LightHouse: generated URL to aggregate attestation", zap.String("url", url))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "LightHouse: failed to create request with context")
	}

	resp, err := n.httpClient.Do(req)
	if err != nil {
		if err == context.DeadlineExceeded || err == context.Canceled {
			return nil, nil
		}

		n.logger.Error("LightHouse: failed to send request", zap.Error(err))
		return nil, errors.Wrap(err, "LightHouse: failed to send request")
	}

	if resp.StatusCode > 299 {
		n.logger.Error("LightHouse: unexpected response code", zap.Int("status_code", resp.StatusCode), zap.String("body", getResponseBodyRaw(resp)))
		return nil, errors.Errorf("LightHouse: unexpected response code %d", resp.StatusCode)
	}

	if resp.Body == nil {
		return nil, errors.New("LightHouse: empty response body")
	}

	defer resp.Body.Close()

	var aggData aggregateModel
	if err := json.NewDecoder(resp.Body).Decode(&aggData); err != nil {
		n.logger.Error("LightHouse: failed to decode response body", zap.Error(err))
		return nil, errors.Wrap(err, "LightHouse: failed to decode response body")
	}

	n.logger.Info("LightHouse: got aggregation", zap.Any("agg_data", aggData))

	return aggData.toProto()
}

// SubmitSignedAggregateSelectionProof verifies given aggregate and proofs and publishes them on appropriate gossipsub topic
func (n *lighthouseHTTP) SubmitSignedAggregateSelectionProof(_ context.Context, signature []byte, message *ethpb.AggregateAttestationAndProof) error {
	reqBody, err := json.Marshal([]interface{}{toSignedAggregateModel(signature, message)})
	if err != nil {
		return errors.Wrap(err, "LightHouse: failed to marshal lighthouse JSON model")
	}

	req, err := http.NewRequest(http.MethodPost, n.addr+"/eth/v1/validator/aggregate_and_proofs", bytes.NewReader(reqBody))
	if err != nil {
		return errors.Wrap(err, "LightHouse: failed to create request with context")
	}

	resp, err := n.httpClient.Do(req)
	if err != nil {
		if err == context.DeadlineExceeded || err == context.Canceled {
			return nil
		}

		n.logger.Error("LightHouse: failed to send request", zap.Error(err), zap.String("req_body", string(reqBody)))
		return errors.Wrap(err, "LightHouse: failed to send request")
	}

	if resp.StatusCode > 299 {
		n.logger.Error("LightHouse: unexpected response code", zap.Int("status_code", resp.StatusCode),
			zap.String("resp_body", getResponseBodyRaw(resp)), zap.String("req_body", string(reqBody)))
		return errors.Errorf("LightHouse: unexpected response code %d", resp.StatusCode)
	}

	return nil
}

// SubnetsSubscribe subscribes on the given subnets
func (n *lighthouseHTTP) SubnetsSubscribe(_ context.Context, subscriptions []beaconchain.SubnetSubscription) error {
	models := make([]subnetSubscriptionModel, len(subscriptions))
	for i, subscription := range subscriptions {
		models[i] = toSubnetSubscriptionModel(subscription)
	}

	reqBody, err := json.Marshal(models)
	if err != nil {
		return errors.Wrap(err, "LightHouse: failed to marshal lighthouse JSON model")
	}

	req, err := http.NewRequest(http.MethodPost, n.addr+"/eth/v1/validator/beacon_committee_subscriptions", bytes.NewReader(reqBody))
	if err != nil {
		return errors.Wrap(err, "LightHouse: failed to create request with context")
	}

	resp, err := n.httpClient.Do(req)
	if err != nil {
		if err == context.DeadlineExceeded || err == context.Canceled {
			return nil
		}

		n.logger.Error("LightHouse: failed to send request", zap.Error(err))
		return errors.Wrap(err, "LightHouse: failed to send request")
	}

	if resp.StatusCode > 299 {
		n.logger.Error("LightHouse: unexpected response code", zap.Int("status_code", resp.StatusCode), zap.String("body", getResponseBodyRaw(resp)))
		return errors.Errorf("LightHouse: unexpected response code %d", resp.StatusCode)
	}

	return nil
}

// DomainData returns domain data by the given request
func (n *lighthouseHTTP) DomainData(ctx context.Context, epoch uint64, domain []byte) ([]byte, error) {
	data, err := n.getGenesisData(ctx)
	if err != nil {
		n.logger.Error("LightHouse: failed to get head fork", zap.Error(err))
		return nil, errors.Wrap(err, "LightHouse: failed to get head fork")
	}

	currentVersion, err := hex.DecodeString(strings.TrimPrefix(data.Data.GenesisForkVersion, "0x"))
	if err != nil {
		n.logger.Error("LightHouse: failed to decode genesis fork version", zap.Error(err), zap.Any("data", data.Data))
		return nil, errors.Wrap(err, "LightHouse: failed to decode genesis fork version")
	}

	genesisValidatorRoot, err := hex.DecodeString(strings.TrimPrefix(data.Data.GenesisValidatorsRoot, "0x"))
	if err != nil {
		n.logger.Error("LightHouse: failed to decode genesis validators root", zap.Error(err), zap.Any("data", data.Data))
		return nil, errors.Wrap(err, "LightHouse: failed to decode genesis validators root")
	}

	dv, err := helpers.Domain(&pb.Fork{CurrentVersion: currentVersion}, epoch, bytesutil.ToBytes4(domain), genesisValidatorRoot)
	if err != nil {
		n.logger.Error("LightHouse: failed to get domain data", zap.Error(err), zap.Any("data", data.Data))
		return nil, errors.Wrap(err, "LightHouse: failed to get domain data")
	}

	n.logger.Info("LightHouse: got domain data", zap.String("domain_data", hex.EncodeToString(dv)))

	return dv, nil
}

// StreamDuties returns client to stream duties
func (n *lighthouseHTTP) StreamDuties(ctx context.Context, pubKeys [][]byte) (ethpb.BeaconNodeValidator_StreamDutiesClient, error) {
	validatorIDs := make([]string, len(pubKeys))
	var wg sync.WaitGroup
	for i, pubKey := range pubKeys {
		wg.Add(1)
		go func(i int, pubKey []byte) {
			defer wg.Done()

			valID, err := n.getValidatorID(ctx, pubKey)
			if err != nil {
				n.logger.Error("failed to get validator ID", zap.Error(err))
			}

			validatorIDs[i] = valID
		}(i, pubKey)
	}
	wg.Wait()

	cl := newLighthouseStreamDuties(ctx, n.logger, n.network, n.addr, validatorIDs)
	return cl, nil
}

// GetGenesis returns genesis data
func (n *lighthouseHTTP) GetGenesis(ctx context.Context) (*ethpb.Genesis, error) {
	genesis, err := n.getGenesisData(ctx)
	if err != nil {
		n.logger.Error("LightHouse: failed to get genesis data", zap.Error(err))
		return nil, errors.Wrap(err, "LightHouse: failed to get genesis data")
	}

	model, err := genesis.toGenesis()
	if err != nil {
		n.logger.Error("LightHouse: failed to build proto model of genesis data", zap.Error(err))
		return nil, errors.Wrap(err, "LightHouse: failed to build proto model of genesis data")
	}

	return model, nil
}

func (n *lighthouseHTTP) getValidatorID(ctx context.Context, pubKey []byte) (string, error) {
	url := n.addr + "/eth/v1/beacon/states/head/validators?id=0x" + hex.EncodeToString(pubKey)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", errors.Wrap(err, "LightHouse: failed to create request with context")
	}

	resp, err := n.httpClient.Do(req)
	if err != nil {
		if err == context.DeadlineExceeded || err == context.Canceled {
			return "", nil
		}

		n.logger.Error("LightHouse: failed to send request to get genesis", zap.Error(err))
		return "", errors.Wrap(err, "LightHouse: failed to send request to get validator")
	}

	if resp.StatusCode > 299 {
		n.logger.Error("LightHouse: unexpected response code", zap.Int("status_code", resp.StatusCode), zap.String("body", getResponseBodyRaw(resp)))
		return "", errors.Errorf("LightHouse: unexpected response code %d", resp.StatusCode)
	}

	if resp.Body == nil {
		return "", errors.New("LightHouse: empty response body")
	}

	defer resp.Body.Close()

	var data struct {
		Data []struct {
			Index string `json:"index"`
		} `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		n.logger.Error("LightHouse: failed to decode response body", zap.Error(err))
		return "", errors.Wrap(err, "LightHouse: failed to decode response body")
	}

	if len(data.Data) == 0 {
		return "", errors.New("validator not found")
	}

	return data.Data[0].Index, nil
}

// getGenesisData loads genesis data from prysm node
func (n *lighthouseHTTP) getGenesisData(ctx context.Context) (genesisModel, error) {
	n.genesisDataLock.Lock()
	defer n.genesisDataLock.Unlock()

	if n.genesisData.Data.GenesisTime != "" {
		return n.genesisData, nil
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, n.addr+"/eth/v1/beacon/genesis", nil)
	if err != nil {
		return n.genesisData, errors.Wrap(err, "LightHouse: failed to create request with context")
	}

	resp, err := n.httpClient.Do(req)
	if err != nil {
		if err == context.DeadlineExceeded || err == context.Canceled {
			return n.genesisData, nil
		}

		n.logger.Error("LightHouse: failed to send request to get genesis", zap.Error(err))
		return n.genesisData, errors.Wrap(err, "LightHouse: failed to send request to get genesis")
	}

	if resp.StatusCode > 299 {
		n.logger.Error("LightHouse: unexpected response code", zap.Int("status_code", resp.StatusCode), zap.String("body", getResponseBodyRaw(resp)))
		return n.genesisData, errors.Errorf("LightHouse: unexpected response code %d", resp.StatusCode)
	}

	if resp.Body == nil {
		return n.genesisData, errors.New("LightHouse: empty response body")
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&n.genesisData); err != nil {
		n.logger.Error("LightHouse: failed to decode response body", zap.Error(err))
		return n.genesisData, errors.Wrap(err, "LightHouse: failed to decode response body")
	}

	return n.genesisData, nil
}

func getResponseBodyRaw(resp *http.Response) string {
	if resp == nil || resp.Body == nil {
		return ""
	}

	bodyRaw, _ := ioutil.ReadAll(resp.Body)
	return string(bodyRaw)
}
