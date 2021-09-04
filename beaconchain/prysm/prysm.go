package prysm

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	types "github.com/prysmaticlabs/eth2-types"
	ethpb "github.com/prysmaticlabs/prysm/proto/prysm/v1alpha1"
	"go.uber.org/zap"

	"github.com/begmaroman/beaconspot/beaconchain"
)

// prysmGRPC implements beaconchain.BeaconChain interface using Prysm beacon node via gRPC
type prysmGRPC struct {
	validatorClient ethpb.BeaconNodeValidatorClient
	nodeClient      ethpb.NodeClient
	logger          *zap.Logger
}

// New is the constructor of prysmGRPC
func New(logger *zap.Logger, validatorClient ethpb.BeaconNodeValidatorClient, nodeClient ethpb.NodeClient) beaconchain.BeaconChain {
	return &prysmGRPC{
		validatorClient: validatorClient,
		nodeClient:      nodeClient,
		logger:          logger,
	}
}

// SubnetsSubscribe subscribes on the given subnets
func (c *prysmGRPC) SubnetsSubscribe(ctx context.Context, subscriptions []beaconchain.SubnetSubscription) error {
	subscribeReq := &ethpb.CommitteeSubnetsSubscribeRequest{
		Slots:        make([]types.Slot, len(subscriptions)),
		CommitteeIds: make([]types.CommitteeIndex, len(subscriptions)),
		IsAggregator: make([]bool, len(subscriptions)),
	}

	for i, subscription := range subscriptions {
		subscribeReq.Slots[i] = subscription.Slot
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
func (c *prysmGRPC) DomainData(ctx context.Context, epoch types.Epoch, domain []byte) ([]byte, error) {
	res, err := c.validatorClient.DomainData(ctx, &ethpb.DomainRequest{
		Epoch:  epoch,
		Domain: domain,
	})
	if err != nil {
		return nil, errors.Wrap(err, "Prysm: failed to get domain data")
	}

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

// GetGenesis returns genesis data
func (c *prysmGRPC) GetGenesis(ctx context.Context) (*ethpb.Genesis, error) {
	res, err := c.nodeClient.GetGenesis(ctx, &empty.Empty{})
	if err != nil {
		return nil, errors.Wrap(err, "Prysm: failed to get genesis")
	}

	return res, nil
}
