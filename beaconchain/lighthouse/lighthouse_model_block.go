package lighthouse

import (
	"encoding/hex"
	"strconv"
	"strings"

	types "github.com/prysmaticlabs/eth2-types"

	"github.com/pkg/errors"
	ethpb "github.com/prysmaticlabs/ethereumapis/eth/v1alpha1"
)

type voluntaryExitsMessageModel struct {
	Epoch          string `json:"epoch"`
	ValidatorIndex string `json:"validator_index"`
}

func (m *voluntaryExitsMessageModel) toProto() (*ethpb.VoluntaryExit, error) {
	epoch, err := strconv.Atoi(m.Epoch)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse epoch '%s'", m.Epoch)
	}

	index, err := strconv.Atoi(m.ValidatorIndex)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse index '%s'", m.ValidatorIndex)
	}

	return &ethpb.VoluntaryExit{
		Epoch:          types.Epoch(epoch),
		ValidatorIndex: uint64(index),
	}, nil
}

func toVoluntaryExitsMessageModel(exit *ethpb.VoluntaryExit) voluntaryExitsMessageModel {
	return voluntaryExitsMessageModel{
		Epoch:          strconv.Itoa(int(exit.GetEpoch())),
		ValidatorIndex: strconv.Itoa(int(exit.GetValidatorIndex())),
	}
}

type voluntaryExitsModel struct {
	Message   voluntaryExitsMessageModel `json:"message"`
	Signature string                     `json:"signature"`
}

func (m *voluntaryExitsModel) toProto() (*ethpb.SignedVoluntaryExit, error) {
	message, err := m.Message.toProto()
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse message")
	}

	signature, err := hex.DecodeString(strings.TrimPrefix(m.Signature, "0x"))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse signature '%s'", m.Signature)
	}

	return &ethpb.SignedVoluntaryExit{
		Exit:      message,
		Signature: signature,
	}, nil
}

func toVoluntaryExitsModel(exit *ethpb.SignedVoluntaryExit) voluntaryExitsModel {
	return voluntaryExitsModel{
		Message:   toVoluntaryExitsMessageModel(exit.GetExit()),
		Signature: "0x" + hex.EncodeToString(exit.GetSignature()),
	}
}

type depositDataModel struct {
	PubKey                string `json:"pubkey"`
	WithdrawalCredentials string `json:"withdrawal_credentials"`
	Amount                string `json:"amount"`
	Signature             string `json:"signature"`
}

func (m *depositDataModel) toProto() (*ethpb.Deposit_Data, error) {
	pubKey, err := hex.DecodeString(strings.TrimPrefix(m.PubKey, "0x"))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse public key '%s'", m.PubKey)
	}

	withdrawalCreds, err := hex.DecodeString(strings.TrimPrefix(m.WithdrawalCredentials, "0x"))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse withdrawal creds '%s'", m.WithdrawalCredentials)
	}

	amount, err := strconv.Atoi(m.Amount)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse amount '%s'", m.Amount)
	}

	signature, err := hex.DecodeString(strings.TrimPrefix(m.Signature, "0x"))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse signature '%s'", m.Signature)
	}

	return &ethpb.Deposit_Data{
		PublicKey:             pubKey,
		WithdrawalCredentials: withdrawalCreds,
		Amount:                uint64(amount),
		Signature:             signature,
	}, nil
}

func toDepositDataModel(data *ethpb.Deposit_Data) depositDataModel {
	return depositDataModel{
		PubKey:                "0x" + hex.EncodeToString(data.GetPublicKey()),
		WithdrawalCredentials: "0x" + hex.EncodeToString(data.GetWithdrawalCredentials()),
		Amount:                strconv.Itoa(int(data.GetAmount())),
		Signature:             "0x" + hex.EncodeToString(data.GetSignature()),
	}
}

type depositModel struct {
	Proof []string         `json:"proof"`
	Data  depositDataModel `json:"data"`
}

func (m *depositModel) toProto() (*ethpb.Deposit, error) {
	var proofs [][]byte
	for _, proof := range m.Proof {
		proofBytes, err := hex.DecodeString(strings.TrimPrefix(proof, "0x"))
		if err != nil {
			return nil, errors.Wrapf(err, "failed to parse proof '%s'", proof)
		}

		proofs = append(proofs, proofBytes)
	}

	data, err := m.Data.toProto()
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse data")
	}

	return &ethpb.Deposit{
		Proof: proofs,
		Data:  data,
	}, nil
}

func toDepositModel(deposit *ethpb.Deposit) depositModel {
	proofs := make([]string, len(deposit.GetProof()))
	for i, proof := range deposit.GetProof() {
		proofs[i] = "0x" + hex.EncodeToString(proof)
	}

	return depositModel{
		Proof: proofs,
		Data:  toDepositDataModel(deposit.GetData()),
	}
}

type indexedAttestationModel struct {
	AttestingIndices []string             `json:"attesting_indices"`
	Signature        string               `json:"signature"`
	Data             attestationDataModel `json:"data"`
}

func (m *indexedAttestationModel) toProto() (*ethpb.IndexedAttestation, error) {
	var attestingIndices []uint64
	for _, idx := range m.AttestingIndices {
		index, err := strconv.Atoi(idx)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to parse attesting index '%s'", idx)
		}

		attestingIndices = append(attestingIndices, uint64(index))
	}

	signature, err := hex.DecodeString(strings.TrimPrefix(m.Signature, "0x"))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse signature '%s'", m.Signature)
	}

	data, err := m.Data.toProto()
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse data")
	}

	return &ethpb.IndexedAttestation{
		AttestingIndices: attestingIndices,
		Data:             data,
		Signature:        signature,
	}, nil
}

func toIndexedAttestationModel(data *ethpb.IndexedAttestation) indexedAttestationModel {
	attestingIndices := make([]string, len(data.GetAttestingIndices()))
	for i, idx := range data.GetAttestingIndices() {
		attestingIndices[i] = strconv.Itoa(int(idx))
	}

	return indexedAttestationModel{
		AttestingIndices: attestingIndices,
		Signature:        "0x" + hex.EncodeToString(data.GetSignature()),
		Data:             toAttestationDataModel(data.GetData()),
	}
}

type attesterSlashingModel struct {
	Attestation1 indexedAttestationModel `json:"attestation_1"`
	Attestation2 indexedAttestationModel `json:"attestation_2"`
}

func (m *attesterSlashingModel) toProto() (*ethpb.AttesterSlashing, error) {
	att1, err := m.Attestation1.toProto()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get attestation 1")
	}

	att2, err := m.Attestation1.toProto()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get attestation 2")
	}

	return &ethpb.AttesterSlashing{
		Attestation_1: att1,
		Attestation_2: att2,
	}, nil
}

func toAttesterSlashingModel(slashing *ethpb.AttesterSlashing) attesterSlashingModel {
	return attesterSlashingModel{
		Attestation1: toIndexedAttestationModel(slashing.GetAttestation_1()),
		Attestation2: toIndexedAttestationModel(slashing.GetAttestation_2()),
	}
}

type blockHeaderModel struct {
	Slot          string `json:"slot"`
	ProposerIndex string `json:"proposer_index"`
	ParentRoot    string `json:"parent_root"`
	StateRoot     string `json:"state_root"`
	BodyRoot      string `json:"body_root"`
}

func (m *blockHeaderModel) toProto() (*ethpb.BeaconBlockHeader, error) {
	slot, err := strconv.Atoi(m.Slot)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse slot '%s'", m.Slot)
	}

	proposerIndex, err := strconv.Atoi(m.ProposerIndex)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse proposer index '%s'", m.ProposerIndex)
	}

	parentRoot, err := hex.DecodeString(strings.TrimPrefix(m.ParentRoot, "0x"))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse parent root '%s'", m.ParentRoot)
	}

	stateRoot, err := hex.DecodeString(strings.TrimPrefix(m.StateRoot, "0x"))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse state root '%s'", m.StateRoot)
	}

	bodyRoot, err := hex.DecodeString(strings.TrimPrefix(m.BodyRoot, "0x"))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse body root '%s'", m.BodyRoot)
	}

	return &ethpb.BeaconBlockHeader{
		Slot:          types.Slot(slot),
		ProposerIndex: uint64(proposerIndex),
		ParentRoot:    parentRoot,
		StateRoot:     stateRoot,
		BodyRoot:      bodyRoot,
	}, nil
}

func toBlockHeaderModel(header *ethpb.BeaconBlockHeader) blockHeaderModel {
	return blockHeaderModel{
		Slot:          strconv.Itoa(int(header.GetSlot())),
		ProposerIndex: strconv.Itoa(int(header.GetProposerIndex())),
		ParentRoot:    "0x" + hex.EncodeToString(header.GetParentRoot()),
		StateRoot:     "0x" + hex.EncodeToString(header.GetStateRoot()),
		BodyRoot:      "0x" + hex.EncodeToString(header.GetBodyRoot()),
	}
}

type signedBeaconBlockHeaderModel struct {
	Message   blockHeaderModel `json:"message"`
	Signature string           `json:"signature"`
}

func (m *signedBeaconBlockHeaderModel) toProto() (*ethpb.SignedBeaconBlockHeader, error) {
	message, err := m.Message.toProto()
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse message")
	}

	signature, err := hex.DecodeString(strings.TrimPrefix(m.Signature, "0x"))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse signature '%s'", m.Signature)
	}

	return &ethpb.SignedBeaconBlockHeader{
		Header:    message,
		Signature: signature,
	}, nil
}

func toSignedBeaconBlockHeaderModel(header *ethpb.SignedBeaconBlockHeader) signedBeaconBlockHeaderModel {
	return signedBeaconBlockHeaderModel{
		Message:   toBlockHeaderModel(header.GetHeader()),
		Signature: "0x" + hex.EncodeToString(header.GetSignature()),
	}
}

type proposerSlashingModel struct {
	SignedHeader1 signedBeaconBlockHeaderModel `json:"signed_header_1"`
	SignedHeader2 signedBeaconBlockHeaderModel `json:"signed_header_2"`
}

func (m *proposerSlashingModel) toProto() (*ethpb.ProposerSlashing, error) {
	head1, err := m.SignedHeader1.toProto()
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse signed header 1")
	}

	head2, err := m.SignedHeader2.toProto()
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse signed header 2")
	}

	return &ethpb.ProposerSlashing{
		Header_1: head1,
		Header_2: head2,
	}, nil
}

func toProposerSlashingModel(slashing *ethpb.ProposerSlashing) proposerSlashingModel {
	return proposerSlashingModel{
		SignedHeader1: toSignedBeaconBlockHeaderModel(slashing.GetHeader_1()),
		SignedHeader2: toSignedBeaconBlockHeaderModel(slashing.GetHeader_2()),
	}
}

type eth1DataModel struct {
	DepositRoot  string `json:"deposit_root"`
	DepositCount string `json:"deposit_count"`
	BlockHash    string `json:"block_hash"`
}

func (m *eth1DataModel) toProto() (*ethpb.Eth1Data, error) {
	depositCount, err := strconv.Atoi(m.DepositCount)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse deposit count '%s'", m.DepositCount)
	}

	depositRoot, err := hex.DecodeString(strings.TrimPrefix(m.DepositRoot, "0x"))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse deposit root '%s'", m.DepositRoot)
	}

	blockHash, err := hex.DecodeString(strings.TrimPrefix(m.BlockHash, "0x"))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse block hash '%s'", m.BlockHash)
	}

	return &ethpb.Eth1Data{
		DepositRoot:  depositRoot,
		DepositCount: uint64(depositCount),
		BlockHash:    blockHash,
	}, nil
}

func toEth1DataModel(data *ethpb.Eth1Data) eth1DataModel {
	return eth1DataModel{
		DepositRoot:  "0x" + hex.EncodeToString(data.GetDepositRoot()),
		DepositCount: strconv.Itoa(int(data.GetDepositCount())),
		BlockHash:    "0x" + hex.EncodeToString(data.GetBlockHash()),
	}
}

type blockBodyModel struct {
	RandaoReveal      string                  `json:"randao_reveal"`
	Eth1Data          eth1DataModel           `json:"eth1_data"`
	Graffiti          string                  `json:"graffiti"`
	ProposerSlashings []proposerSlashingModel `json:"proposer_slashings"`
	AttesterSlashings []attesterSlashingModel `json:"attester_slashings"`
	Attestations      []attestationModel      `json:"attestations"`
	Deposits          []depositModel          `json:"deposits"`
	VoluntaryExits    []voluntaryExitsModel   `json:"voluntary_exits"`
}

func (m *blockBodyModel) toProto() (*ethpb.BeaconBlockBody, error) {
	randaoReveal, err := hex.DecodeString(strings.TrimPrefix(m.RandaoReveal, "0x"))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse randao reveal '%s'", m.RandaoReveal)
	}

	eth1Data, err := m.Eth1Data.toProto()
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse eth1 data")
	}

	graffiti, err := hex.DecodeString(strings.TrimPrefix(m.Graffiti, "0x"))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse graffiti '%s'", m.Graffiti)
	}

	var proposerSlashings []*ethpb.ProposerSlashing
	for _, proposerSlashing := range m.ProposerSlashings {
		proposerSlashingProto, err := proposerSlashing.toProto()
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse proposer slashing")
		}

		proposerSlashings = append(proposerSlashings, proposerSlashingProto)
	}

	var attesterSlashings []*ethpb.AttesterSlashing
	for _, attesterSlashing := range m.AttesterSlashings {
		attesterSlashingProto, err := attesterSlashing.toProto()
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse attester slashing")
		}

		attesterSlashings = append(attesterSlashings, attesterSlashingProto)
	}

	var attestations []*ethpb.Attestation
	for _, attestation := range m.Attestations {
		attestationProto, err := attestation.toProto()
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse attestation")
		}

		attestations = append(attestations, attestationProto)
	}

	var deposits []*ethpb.Deposit
	for _, deposit := range m.Deposits {
		depositProto, err := deposit.toProto()
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse deposit")
		}

		deposits = append(deposits, depositProto)
	}

	var voluntaryExits []*ethpb.SignedVoluntaryExit
	for _, voluntaryExit := range m.VoluntaryExits {
		voluntaryExitProto, err := voluntaryExit.toProto()
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse voluntary exit")
		}

		voluntaryExits = append(voluntaryExits, voluntaryExitProto)
	}

	return &ethpb.BeaconBlockBody{
		RandaoReveal:      randaoReveal,
		Eth1Data:          eth1Data,
		Graffiti:          graffiti,
		ProposerSlashings: proposerSlashings,
		AttesterSlashings: attesterSlashings,
		Attestations:      attestations,
		Deposits:          deposits,
		VoluntaryExits:    voluntaryExits,
	}, nil
}

func toBlockBodyModel(body *ethpb.BeaconBlockBody) blockBodyModel {
	proposerSlashings := make([]proposerSlashingModel, len(body.GetProposerSlashings()))
	for i, proposerSlashing := range body.GetProposerSlashings() {
		proposerSlashings[i] = toProposerSlashingModel(proposerSlashing)
	}

	attesterSlashings := make([]attesterSlashingModel, len(body.GetAttesterSlashings()))
	for i, attesterSlashing := range body.GetAttesterSlashings() {
		attesterSlashings[i] = toAttesterSlashingModel(attesterSlashing)
	}

	attestations := make([]attestationModel, len(body.GetAttestations()))
	for i, attestation := range body.GetAttestations() {
		attestations[i] = toAttestationModel(attestation)
	}

	deposits := make([]depositModel, len(body.GetDeposits()))
	for i, deposit := range body.GetDeposits() {
		deposits[i] = toDepositModel(deposit)
	}

	voluntaryExits := make([]voluntaryExitsModel, len(body.GetVoluntaryExits()))
	for i, exit := range body.GetVoluntaryExits() {
		voluntaryExits[i] = toVoluntaryExitsModel(exit)
	}

	return blockBodyModel{
		RandaoReveal:      "0x" + hex.EncodeToString(body.GetRandaoReveal()),
		Eth1Data:          toEth1DataModel(body.GetEth1Data()),
		Graffiti:          "0x" + hex.EncodeToString(body.GetGraffiti()),
		ProposerSlashings: proposerSlashings,
		AttesterSlashings: attesterSlashings,
		Attestations:      attestations,
		Deposits:          deposits,
		VoluntaryExits:    voluntaryExits,
	}
}

type blockModel struct {
	Slot          string         `json:"slot"`
	ProposerIndex string         `json:"proposer_index"`
	ParentRoot    string         `json:"parent_root"`
	StateRoot     string         `json:"state_root"`
	Body          blockBodyModel `json:"body"`
}

func (m *blockModel) toProto() (*ethpb.BeaconBlock, error) {
	slot, err := strconv.Atoi(m.Slot)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse slot '%s'", m.Slot)
	}

	proposerIndex, err := strconv.Atoi(m.ProposerIndex)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse proposer index '%s'", m.ProposerIndex)
	}

	parentRoot, err := hex.DecodeString(strings.TrimPrefix(m.ParentRoot, "0x"))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse parent root '%s'", m.ParentRoot)
	}

	stateRoot, err := hex.DecodeString(strings.TrimPrefix(m.StateRoot, "0x"))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse state root '%s'", m.StateRoot)
	}

	body, err := m.Body.toProto()
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse body")
	}

	return &ethpb.BeaconBlock{
		Slot:          types.Slot(slot),
		ProposerIndex: uint64(proposerIndex),
		ParentRoot:    parentRoot,
		StateRoot:     stateRoot,
		Body:          body,
	}, nil
}

func toBlockModel(block *ethpb.BeaconBlock) blockModel {
	return blockModel{
		Slot:          strconv.Itoa(int(block.GetSlot())),
		ProposerIndex: strconv.Itoa(int(block.GetProposerIndex())),
		ParentRoot:    "0x" + hex.EncodeToString(block.GetParentRoot()),
		StateRoot:     "0x" + hex.EncodeToString(block.GetStateRoot()),
		Body:          toBlockBodyModel(block.GetBody()),
	}
}

type blockReqModel struct {
	Data blockModel `json:"data"`
}

type blockSubmitModel struct {
	Message   blockModel `json:"message"`
	Signature string     `json:"signature"`
}

func toBlockSubmitModel(signature []byte, block *ethpb.BeaconBlock) blockSubmitModel {
	return blockSubmitModel{
		Message:   toBlockModel(block),
		Signature: "0x" + hex.EncodeToString(signature),
	}
}
