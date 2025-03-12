package messages

import (
	"encoding/json"
)

type SubmitBidMessage struct {
	Bidder       string   `json:"bidder"`
	AuctionId    string   `json:"auctionId"`
	Round        int      `json:"round"`
	GasPrice     int      `json:"gasPrice"`
	Transactions []string `json:"transactions"`
}

func (m *SubmitBidMessage) MessageType() MessageType {
	return SubmitBid
}

func (m *SubmitBidMessage) Marshal() ([]byte, error) {
	return json.Marshal(m)
}
