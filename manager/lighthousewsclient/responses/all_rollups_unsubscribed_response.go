package responses

import (
	"encoding/json"
)

type AllRollupsUnsubscribedResponse struct {
	BidderAddress string `json:"bidderAddress"`
}

func (r *AllRollupsUnsubscribedResponse) Unmarshal(data []byte) error {
	if err := json.Unmarshal(data, r); err != nil {
		return err
	}
	return nil
}
