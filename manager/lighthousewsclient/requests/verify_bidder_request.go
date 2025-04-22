package requests

import (
	"encoding/json"
)

type VerifyBidderRequest struct {
	BidderAddress string `json:"bidderAddress"`
	Signature     []byte `json:"signature"`
}

func (r *VerifyBidderRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}
