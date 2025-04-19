package requests

import "encoding/json"

type UnsubscribeAllRollupsRequest struct {
	BidderAddress *string `json:"bidderAddress"`
}

func (r *UnsubscribeAllRollupsRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}
