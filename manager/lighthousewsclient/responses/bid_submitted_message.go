package responses

import "encoding/json"

type BidSubmittedResponse struct {
	BidderAddress *string  `json:"bidderAddress"`
	AuctionId     *string  `json:"auctionId"`
	Round         *int     `json:"round"`
	BidPrice      *int     `json:"bidPrice"`
	Transactions  []string `json:"transactions"`
}

func (r *BidSubmittedResponse) Unmarshal(data []byte) error {
	if err := json.Unmarshal(data, r); err != nil {
		return err
	}
	return nil
}
