package requests

import (
	"encoding/json"
)

type SubmitBidRequest struct {
	BidderAddress   string   `json:"bidderAddress"`
	AuctionId       string   `json:"auctionId"`
	BidPrice        string   `json:"bidPrice"`
	RawTransactions [][]byte `json:"rawTransactions"`
}

func (r *SubmitBidRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}
