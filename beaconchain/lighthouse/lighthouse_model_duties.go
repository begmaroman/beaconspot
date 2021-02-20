package lighthouse

import (
	"encoding/hex"
	"strconv"
	"strings"

	types "github.com/prysmaticlabs/eth2-types"

	"github.com/pkg/errors"
	ethpb "github.com/prysmaticlabs/ethereumapis/eth/v1alpha1"
)

type dutyDataModel struct {
	PubKey                  string `json:"pubkey"`
	ValidatorIndex          string `json:"validator_index"`
	CommitteeIndex          string `json:"committee_index"`
	CommitteeLength         string `json:"committee_length"`
	CommitteesAtSlot        string `json:"committees_at_slot"`
	ValidatorCommitteeIndex string `json:"validator_committee_index"`
	Slot                    string `json:"slot"`
}

func (m *dutyDataModel) toProposerDuties() (*ethpb.DutiesResponse_Duty, error) {
	slot, err := strconv.Atoi(m.Slot)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse slot")
	}

	pubKey, err := hex.DecodeString(strings.TrimPrefix(m.PubKey, "0x"))
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse public key")
	}

	validatorIndex, err := strconv.Atoi(m.ValidatorIndex)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse validator index")
	}

	return &ethpb.DutiesResponse_Duty{
		PublicKey:      pubKey,
		ValidatorIndex: uint64(validatorIndex),
		ProposerSlots:  []types.Slot{types.Slot(slot)},
		Status:         ethpb.ValidatorStatus_ACTIVE, // TODO: Fill real status
	}, nil
}

func (m *dutyDataModel) toAttesterDuties() (*ethpb.DutiesResponse_Duty, error) {
	slot, err := strconv.Atoi(m.Slot)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse slot")
	}

	pubKey, err := hex.DecodeString(strings.TrimPrefix(m.PubKey, "0x"))
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse public key")
	}

	validatorIndex, err := strconv.Atoi(m.ValidatorIndex)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse validator index")
	}

	committeeIndex, err := strconv.Atoi(m.CommitteeIndex)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse committee index")
	}

	committeeLen, err := strconv.Atoi(m.CommitteeLength)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse committee length")
	}

	return &ethpb.DutiesResponse_Duty{
		PublicKey:      pubKey,
		ValidatorIndex: uint64(validatorIndex),
		CommitteeIndex: uint64(committeeIndex),
		Committee:      make([]uint64, committeeLen), // TODO: Implement real committee
		AttesterSlot:   types.Slot(slot),
		Status:         ethpb.ValidatorStatus_ACTIVE, // TODO: Fill real status
	}, nil
}

type dutyModel struct {
	DependentRoot string          `json:"dependent_root"`
	Data          []dutyDataModel `json:"data"`
}
