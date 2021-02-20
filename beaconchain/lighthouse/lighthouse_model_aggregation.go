package lighthouse

import (
	"encoding/hex"
	"strconv"

	"github.com/pkg/errors"
	ethpb "github.com/prysmaticlabs/ethereumapis/eth/v1alpha1"
)

type aggregateModel struct {
	Data attestationModel `json:"data"`
}

func (m *aggregateModel) toProto() (*ethpb.AggregateAttestationAndProof, error) {
	data, err := m.Data.toProto()
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse aggregation data")
	}

	return &ethpb.AggregateAttestationAndProof{
		Aggregate: data,
	}, nil
}

type signedAggregateMessageModel struct {
	AggregatorIndex string           `json:"aggregator_index"`
	SelectionProof  string           `json:"selection_proof"`
	Aggregate       attestationModel `json:"aggregate"`
}

func toSignedAggregateMessageModel(message *ethpb.AggregateAttestationAndProof) signedAggregateMessageModel {
	return signedAggregateMessageModel{
		AggregatorIndex: strconv.Itoa(int(message.GetAggregatorIndex())),
		SelectionProof:  "0x" + hex.EncodeToString(message.GetSelectionProof()),
		Aggregate:       toAttestationModel(message.GetAggregate()),
	}
}

type signedAggregateModel struct {
	Message   signedAggregateMessageModel `json:"message"`
	Signature string                      `json:"signature"`
}

func toSignedAggregateModel(sig []byte, message *ethpb.AggregateAttestationAndProof) signedAggregateModel {
	return signedAggregateModel{
		Message:   toSignedAggregateMessageModel(message),
		Signature: "0x" + hex.EncodeToString(sig),
	}
}
