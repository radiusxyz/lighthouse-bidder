package messages

import "encoding/json"

type RegisterBidderMessage struct {
	BidderAddress string `json:"bidderAddress"`
	RollupId      string `json:"rollupId"`
}

func (m *RegisterBidderMessage) MessageType() MessageType {
	return RegisterBidder
}

func (m *RegisterBidderMessage) Marshal() ([]byte, error) {
	return json.Marshal(m)
}
