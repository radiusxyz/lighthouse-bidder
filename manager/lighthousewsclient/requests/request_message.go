package requests

import (
	"encoding/json"
)

type RequestType string

const (
	RegisterBidder RequestType = "RegisterBidder"
	SubmitBid      RequestType = "SubmitBid"
)

type RequestParams interface {
	Marshal() ([]byte, error)
}

type RequestMessage struct {
	Id     string          `json:"id"`
	Method RequestType     `json:"method"`
	Params json.RawMessage `json:"params"`
}

func (r *RequestMessage) Marshal() ([]byte, error) {
	return json.Marshal(r)
}
