package responses

import "encoding/json"

type BidderVerifiedResponse struct {
	BidderAddress *string `json:"bidderAddress"`
}

func (r *BidderVerifiedResponse) Unmarshal(data []byte) error {
	if err := json.Unmarshal(data, r); err != nil {
		return err
	}
	return nil
}
