package events

import (
	"encoding/json"
)

type EventType string

var (
	AuctionStarted EventType = "AuctionStarted"
	Tob            EventType = "Tob"
)

type EventMessage struct {
	EventType EventType       `json:"eventType"`
	Payload   json.RawMessage `json:"payload"`
}

func (m *EventMessage) Marshal() ([]byte, error) {
	return json.Marshal(m)
}
