package messages

import (
	"encoding/json"
)

type RollupRegisteredMessage struct {
	Status *Status `json:"status"`
}

func (m *RollupRegisteredMessage) MessageType() MessageType {
	return RollupRegistered
}

func (m *RollupRegisteredMessage) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

func (m *RollupRegisteredMessage) Validate() error {
	return validateRequiredFields(map[string]any{
		"status": m.Status,
	})
}
