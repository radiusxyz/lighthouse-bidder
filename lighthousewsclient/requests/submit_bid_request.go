package requests

import (
	"encoding/json"
)

type SubmitBidRequest struct {
	BidderAddress string   `json:"bidderAddress"`
	AuctionId     string   `json:"auctionId"`
	Round         int      `json:"round"`
	BidPrice      string   `json:"bidPrice"`
	Transactions  []string `json:"transactions"`
}

func (r *SubmitBidRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}
