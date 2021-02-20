package lighthouse

import (
	"encoding/hex"
	"strconv"
	"strings"

	types "github.com/prysmaticlabs/eth2-types"

	"github.com/pkg/errors"
	ethpb "github.com/prysmaticlabs/ethereumapis/eth/v1alpha1"
)

type checkpointModel struct {
	Epoch string `json:"epoch"`
	Root  string `json:"root"`
}

func (m *checkpointModel) toProto() (*ethpb.Checkpoint, error) {
	epoch, err := strconv.Atoi(m.Epoch)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse epoch '%s'", m.Epoch)
	}

	root, err := hex.DecodeString(strings.TrimPrefix(m.Root, "0x"))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse root '%s'", m.Epoch)
	}

	return &ethpb.Checkpoint{
		Epoch: types.Epoch(epoch),
		Root:  root,
	}, nil
}

func toCheckpointModel(checkpoint *ethpb.Checkpoint) checkpointModel {
	return checkpointModel{
		Epoch: strconv.Itoa(int(checkpoint.GetEpoch())),
		Root:  "0x" + hex.EncodeToString(checkpoint.GetRoot()),
	}
}

type attestationDataModel struct {
	BeaconBlockRoot string          `json:"beacon_block_root"`
	Index           string          `json:"index"`
	Slot            string          `json:"slot"`
	Source          checkpointModel `json:"source"`
	Target          checkpointModel `json:"target"`
}

func (m *attestationDataModel) toProto() (*ethpb.AttestationData, error) {
	dataSlot, err := strconv.Atoi(m.Slot)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse slot '%s'", m.Slot)
	}

	dataIndex, err := strconv.Atoi(m.Index)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse index '%s'", m.Index)
	}

	dataBeaconBlockRoot, err := hex.DecodeString(strings.TrimPrefix(m.BeaconBlockRoot, "0x"))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse beacon block root '%s'", m.BeaconBlockRoot)
	}

	source, err := m.Source.toProto()
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse source")
	}

	target, err := m.Target.toProto()
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse source")
	}

	return &ethpb.AttestationData{
		Slot:            types.Slot(dataSlot),
		CommitteeIndex:  uint64(dataIndex),
		BeaconBlockRoot: dataBeaconBlockRoot,
		Source:          source,
		Target:          target,
	}, nil
}

func toAttestationDataModel(data *ethpb.AttestationData) attestationDataModel {
	return attestationDataModel{
		BeaconBlockRoot: "0x" + hex.EncodeToString(data.GetBeaconBlockRoot()),
		Index:           strconv.Itoa(int(data.GetCommitteeIndex())),
		Slot:            strconv.Itoa(int(data.GetSlot())),
		Source:          toCheckpointModel(data.GetSource()),
		Target:          toCheckpointModel(data.GetTarget()),
	}
}

type attestationModel struct {
	AggregationBits string               `json:"aggregation_bits"`
	Signature       string               `json:"signature"`
	Data            attestationDataModel `json:"data"`
}

func (m *attestationModel) toProto() (*ethpb.Attestation, error) {
	data, err := m.Data.toProto()
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse data")
	}

	signature, err := hex.DecodeString(strings.TrimPrefix(m.Signature, "0x"))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse signature '%s'", m.Signature)
	}

	aggregationBits, err := hex.DecodeString(strings.TrimPrefix(m.AggregationBits, "0x"))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse aggregation bits '%s'", m.AggregationBits)
	}

	return &ethpb.Attestation{
		AggregationBits: aggregationBits,
		Data:            data,
		Signature:       signature,
	}, nil
}

func toAttestationModel(att *ethpb.Attestation) attestationModel {
	return attestationModel{
		AggregationBits: "0x" + hex.EncodeToString(att.GetAggregationBits()),
		Signature:       "0x" + hex.EncodeToString(att.GetSignature()),
		Data:            toAttestationDataModel(att.GetData()),
	}
}
