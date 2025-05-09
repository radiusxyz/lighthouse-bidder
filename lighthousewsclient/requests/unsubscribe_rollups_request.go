package requests

import "encoding/json"

type UnsubscribeRollupsRequest struct {
	BidderAddress string   `json:"bidderAddress"`
	RollupIds     []string `json:"rollupIds"`
}

func (r *UnsubscribeRollupsRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}
