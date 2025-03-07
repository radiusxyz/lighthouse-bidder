package messages

import (
	"encoding/json"
	"github.com/radiusxyz/lighthouse-bidder/types"
)

type RoundStartedMessage struct {
	AuctionId    string               `json:"auctionId"`
	Round        int                  `json:"round"`
	Transactions []*types.Transaction `json:"transactions"`
}

func NewRoundStartedMessage(auctionId string, round int, transactions []*types.Transaction) *RoundStartedMessage {
	return &RoundStartedMessage{
		AuctionId:    auctionId,
		Round:        round,
		Transactions: transactions,
	}
}

func (m *RoundStartedMessage) MessageType() MessageType {
	return RoundStarted
}

func (m *RoundStartedMessage) Marshal() ([]byte, error) {
	return json.Marshal(m)
}
