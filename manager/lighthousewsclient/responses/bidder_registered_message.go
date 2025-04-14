package responses

import "encoding/json"

type BidderRegisteredResponse struct {
	Status Status `json:"status"`
}

func (r *BidderRegisteredResponse) Unmarshal(data []byte) error {
	if err := json.Unmarshal(data, r); err != nil {
		return err
	}
	return nil
}
