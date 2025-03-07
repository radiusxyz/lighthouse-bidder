package messages

import "encoding/json"

type CreateAuctionMessage struct {
	RollupId    string `json:"rollupId"`
	BlockNumber uint64 `json:"blockNumber"`
	BlockTime   int    `json:"blockTime"`
}

func (m *CreateAuctionMessage) MessageType() MessageType {
	return CreateAuction
}

func (m *CreateAuctionMessage) Marshal() ([]byte, error) {
	return json.Marshal(m)
}
