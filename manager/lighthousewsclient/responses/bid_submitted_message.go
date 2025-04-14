package responses

import "encoding/json"

type BidSubmittedResponse struct {
	Status Status `json:"status"`
}

func (r *BidSubmittedResponse) Unmarshal(data []byte) error {
	if err := json.Unmarshal(data, r); err != nil {
		return err
	}
	return nil
}
