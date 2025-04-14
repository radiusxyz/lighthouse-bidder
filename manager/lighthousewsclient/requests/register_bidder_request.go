package requests

import (
	"encoding/json"
)

type RegisterBidderRequest struct {
	BidderAddress string `json:"bidderAddress"`
	RollupId      string `json:"rollupId"`
}

func (r *RegisterBidderRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}
