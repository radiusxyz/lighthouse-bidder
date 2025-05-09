package requests

import (
	"encoding/json"
)

type RequestType string

const (
	VerifyBidder          RequestType = "VerifyBidder"
	SubmitBid             RequestType = "SubmitBid"
	SubscribeRollups      RequestType = "SubscribeRollups"
	UnsubscribeRollups    RequestType = "UnsubscribeRollups"
	UnSubscribeAllRollups RequestType = "UnSubscribeAllRollups"
)

type RequestParams interface {
	Marshal() ([]byte, error)
}

type RequestMessage struct {
	Id          string          `json:"id"`
	RequestType RequestType     `json:"requestType"`
	Payload     json.RawMessage `json:"payload"`
}

func (r *RequestMessage) Marshal() ([]byte, error) {
	return json.Marshal(r)
}
