package messages

import "encoding/json"

type RegisterRollupMessage struct {
	Address  string `json:"address"`
	RollupId string `json:"rollupId"`
	RpcUrl   string `json:"rpcUrl"`
	ChainID  string `json:"chainId"`
	SbbUrl   string `json:"sbbUrl"` // Todo SbbUrl -> ClusterId
}

func (m *RegisterRollupMessage) MessageType() MessageType {
	return RegisterRollup
}

func (m *RegisterRollupMessage) Marshal() ([]byte, error) {
	return json.Marshal(m)
}
