package messages

import (
	"encoding/json"
	"github.com/radiusxyz/lighthouse-bidder/types"
)

type RoundStartedMessage struct {
	AuctionId             *string              `json:"auctionId"`
	Round                 *int                 `json:"round"`
	ConfirmedTransactions []*types.Transaction `json:"confirmedTransactions"`
}

func (m *RoundStartedMessage) MessageType() MessageType {
	return RoundStarted
}

func (m *RoundStartedMessage) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

func (m *RoundStartedMessage) Validate() error {
	return validateRequiredFields(map[string]any{
		"auctionId":             m.AuctionId,
		"round":                 m.Round,
		"confirmedTransactions": m.ConfirmedTransactions,
	})
}
