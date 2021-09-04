package lighthouse

import (
	"strconv"

	"github.com/begmaroman/beaconspot/beaconchain"
)

type subnetSubscriptionModel struct {
	ValidatorIndex   string `json:"validator_index"`
	CommitteeIndex   string `json:"committee_index"`
	CommitteesAtSlot string `json:"committees_at_slot"`
	Slot             string `json:"slot"`
	IsAggregator     bool   `json:"is_aggregator"`
}

func toSubnetSubscriptionModel(sub beaconchain.SubnetSubscription) subnetSubscriptionModel {
	return subnetSubscriptionModel{
		ValidatorIndex:   strconv.Itoa(int(sub.ValidatorIndex)),
		CommitteeIndex:   strconv.Itoa(int(sub.CommitteeIndex)),
		CommitteesAtSlot: strconv.Itoa(int(sub.CommitteesAtSlot)),
		Slot:             strconv.Itoa(int(sub.Slot)),
		IsAggregator:     sub.IsAggregator,
	}
}
