package lighthouse

import (
	"encoding/hex"
	"strings"

	"github.com/pkg/errors"
	ethpb "github.com/prysmaticlabs/prysm/proto/prysm/v1alpha1"
)

type genesisDataModel struct {
	GenesisTime           string `json:"genesis_time"`
	GenesisValidatorsRoot string `json:"genesis_validators_root"`
	GenesisForkVersion    string `json:"genesis_fork_version"`
}

type genesisModel struct {
	Data genesisDataModel `json:"data"`
}

func (m *genesisModel) toGenesis() (*ethpb.Genesis, error) {
	genesisValidatorsRoot, err := hex.DecodeString(strings.TrimPrefix(m.Data.GenesisValidatorsRoot, "0x"))
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse public key")
	}

	return &ethpb.Genesis{
		GenesisTime:            nil, // TODO: Fill this field
		DepositContractAddress: nil, // TODO: Fill this field
		GenesisValidatorsRoot:  genesisValidatorsRoot,
	}, nil
}
