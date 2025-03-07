package messages

import "encoding/json"

type AuctionCreatedMessage struct {
	AuctionId *string `json:"auctionId"`
	Status    *Status `json:"status"`
}

func (m *AuctionCreatedMessage) MessageType() MessageType {
	return AuctionCreated
}

func (m *AuctionCreatedMessage) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

func (m *AuctionCreatedMessage) Validate() error {
	return validateRequiredFields(map[string]any{
		"auctionId": m.AuctionId,
		"status":    m.Status,
	})
}
