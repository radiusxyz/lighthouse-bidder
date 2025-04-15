package responses

import "encoding/json"

type BidderRegisteredResponse struct {
	BidderAddress *string `json:"bidderAddress"`
	RollupId      *string `json:"rollupId"`
}

func (r *BidderRegisteredResponse) Unmarshal(data []byte) error {
	if err := json.Unmarshal(data, r); err != nil {
		return err
	}
	return nil
}
