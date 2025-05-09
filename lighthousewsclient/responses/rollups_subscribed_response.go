package responses

import (
	"encoding/json"
)

type RollupsSubscribedResponse struct {
	BidderAddress string   `json:"bidderAddress"`
	RollupIds     []string `json:"rollupIds"`
}

func (r *RollupsSubscribedResponse) Unmarshal(data []byte) error {
	if err := json.Unmarshal(data, r); err != nil {
		return err
	}
	return nil
}
