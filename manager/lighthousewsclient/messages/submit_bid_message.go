package messages

import (
	"encoding/json"
)

type SubmitBidMessage struct {
	BidderAddress string   `json:"bidderAddress"`
	AuctionId     string   `json:"auctionId"`
	Round         int      `json:"round"`
	BidPrice      int      `json:"bidPrice"`
	Transactions  []string `json:"transactions"`
}

func (m *SubmitBidMessage) MessageType() MessageType {
	return SubmitBid
}

func (m *SubmitBidMessage) Marshal() ([]byte, error) {
	return json.Marshal(m)
}
