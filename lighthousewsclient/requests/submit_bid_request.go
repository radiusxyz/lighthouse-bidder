package requests

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
)

type SubmitBidRequest struct {
	BidderAddress   common.Address `json:"bidderAddress"`
	AuctionId       string         `json:"auctionId"`
	BidPrice        string         `json:"bidPrice"`
	RawTransactions [][]byte       `json:"rawTransactions"`
}

func (r *SubmitBidRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func NewSubmitBidRequest() *SubmitBidRequest {
	return &SubmitBidRequest{}
}
