package messages

import (
	"encoding/json"
	"github.com/radiusxyz/lighthouse-bidder/types"
)

type TobMessage struct {
	AuctionId             string               `json:"auctionId"`
	ConfirmedTransactions []*types.Transaction `json:"confirmedTransactions"`
}

func (m *TobMessage) MessageType() MessageType {
	return Tob
}

func (m *TobMessage) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

func (m *TobMessage) Validate() error {
	return validateRequiredFields(map[string]any{
		"auctionId":             m.AuctionId,
		"confirmedTransactions": m.ConfirmedTransactions,
	})
}
