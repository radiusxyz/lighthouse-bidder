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

type ResponsePayload interface {
	Unmarshal(data []byte) error
}

type ResponseMessage struct {
	Id           string          `json:"id"`
	ResponseType ResponseType    `json:"responseType"`
	Status       int             `json:"status"`
	Payload      json.RawMessage `json:"payload"`
	Error        *ErrorMessage   `json:"error"`
}

type ErrorMessage struct {
	Message string `json:"message"`
}
