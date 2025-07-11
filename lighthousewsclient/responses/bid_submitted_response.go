package responses

import "encoding/json"

type BidSubmittedResponse struct {
	BidderAddress   *string  `json:"bidderAddress"`
	AuctionId       *string  `json:"auctionId"`
	BidPrice        *string  `json:"bidPrice"`
	RawTransactions [][]byte `json:"rawTransactions"`
}

func (r *BidSubmittedResponse) Unmarshal(data []byte) error {
	if err := json.Unmarshal(data, r); err != nil {
		return err
	}
	return nil
}
