package responses

import (
	"encoding/json"
)

type Status int

type ResponseType string

const (
	BidderRegistered ResponseType = "BidderRegistered"
	BidSubmitted     ResponseType = "BidSubmitted"
)

type ResponseMessage struct {
	Id           string          `json:"id"`
	ResponseType ResponseType    `json:"responseType"`
	Status       int             `json:"status"`
	Result       json.RawMessage `json:"result"`
	Error        *ErrorMessage   `json:"error"`
}

type ErrorMessage struct {
	Message string `json:"message"`
}

func (m *ResponseMessage) Marshal() ([]byte, error) {
	return json.Marshal(m)
}
