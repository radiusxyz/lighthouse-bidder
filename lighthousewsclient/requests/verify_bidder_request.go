package requests

import (
	"encoding/json"
)

type VerifyBidderRequest struct {
	BidderAddress string `json:"bidderAddress"`
	Timestamp     uint64 `json:"timestamp"`
	Signature     []byte `json:"signature"`
}

func (r *VerifyBidderRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}
