package lighthouse

type genesisDataModel struct {
	GenesisTime           string `json:"genesis_time"`
	GenesisValidatorsRoot string `json:"genesis_validators_root"`
	GenesisForkVersion    string `json:"genesis_fork_version"`
}

type genesisModel struct {
	Data genesisDataModel `json:"data"`
}
