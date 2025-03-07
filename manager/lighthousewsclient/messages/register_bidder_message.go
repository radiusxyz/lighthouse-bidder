package messages

import "encoding/json"

type RegisterBidderMessage struct {
	Bidder   string `json:"bidder"`
	RollupId string `json:"rollupId"`
}

func (m *RegisterBidderMessage) MessageType() MessageType {
	return RegisterBidder
}

func (m *RegisterBidderMessage) Marshal() ([]byte, error) {
	return json.Marshal(m)
}
