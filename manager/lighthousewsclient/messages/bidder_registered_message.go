package messages

import "encoding/json"

type BidderRegisteredMessage struct {
	Status Status `json:"status"`
}

func (m *BidderRegisteredMessage) MessageType() MessageType {
	return BidderRegistered
}

func (m *BidderRegisteredMessage) Marshal() ([]byte, error) {
	return json.Marshal(m)
}
