package messages

import (
	"encoding/json"
	"github.com/radiusxyz/lighthouse-bidder/types"
)

type TobMessage struct {
	AuctionId    string               `json:"auctionId"`
	Transactions []*types.Transaction `json:"transactions"`
}

func NewTobMessage(auctionId string, transactions []*types.Transaction) *TobMessage {
	return &TobMessage{
		AuctionId:    auctionId,
		Transactions: transactions,
	}
}

func (m *TobMessage) MessageType() MessageType {
	return Tob
}

func (m *TobMessage) Marshal() ([]byte, error) {
	return json.Marshal(m)
}
