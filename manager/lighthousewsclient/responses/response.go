package responses

import (
	"encoding/json"
)

type ResponseType string

const (
	BidderVerified         ResponseType = "BidderVerified"
	BidSubmitted           ResponseType = "BidSubmitted"
	RollupsSubscribed      ResponseType = "RollupsSubscribed"
	RollupsUnsubscribed    ResponseType = "RollupsUnsubscribed"
	AllRollupsUnsubscribed ResponseType = "AllRollupsUnsubscribed"
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
