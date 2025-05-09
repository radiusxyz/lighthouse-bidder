package responses

import (
	"encoding/json"
)

type RollupsUnsubscribedResponse struct {
	BidderAddress string   `json:"bidderAddress"`
	RollupIds     []string `json:"rollupIds"`
}

func (r *RollupsUnsubscribedResponse) Unmarshal(data []byte) error {
	if err := json.Unmarshal(data, r); err != nil {
		return err
	}
	return nil
}
