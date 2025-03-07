package messages

import (
	"encoding/json"
	"github.com/radiusxyz/lighthouse-bidder/types"
)

type SubmitBidMessage struct {
	Bidder       string               `json:"bidder"`
	AuctionID    string               `json:"auctionId"`
	Round        int                  `json:"round"`
	GasPrice     int                  `json:"gasPrice"`
	Transactions []*types.Transaction `json:"transactions"`
}

func (m *SubmitBidMessage) MessageType() MessageType {
	return SubmitBid
}

func (m *SubmitBidMessage) Marshal() ([]byte, error) {
	return json.Marshal(m)
}
