package requests

import "encoding/json"

type SubscribeRollupsRequest struct {
	BidderAddress string   `json:"bidderAddress"`
	RollupIds     []string `json:"rollupIds"`
}

func (r *SubscribeRollupsRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}
