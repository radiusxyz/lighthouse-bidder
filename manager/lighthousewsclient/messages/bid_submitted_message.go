package messages

import "encoding/json"

type BidSubmittedMessage struct {
	Status Status `json:"status"`
}

func (m *BidSubmittedMessage) MessageType() MessageType {
	return BidSubmitted
}

func (m *BidSubmittedMessage) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

func (m *BidSubmittedMessage) Validate() error {
	return validateRequiredFields(map[string]any{
		"status": m.Status,
	})
}
