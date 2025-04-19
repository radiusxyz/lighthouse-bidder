package requests

import (
	"encoding/json"
)

type VerifyBidderRequest struct {
	BidderAddress string `json:"bidderAddress"`
}

func (r *VerifyBidderRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}
