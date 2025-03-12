package messages

import "fmt"

type Status int

const (
	Success Status = iota
	Failure
)

type MessageType string

const (
	RegisterRollup   MessageType = "RegisterRollup"
	RegisterBidder   MessageType = "RegisterBidder"
	RollupRegistered MessageType = "RollupRegistered"
	BidderRegistered MessageType = "BidderRegistered"
	CreateAuction    MessageType = "CreateAuction"
	AuctionCreated   MessageType = "AuctionCreated"
	RoundStarted     MessageType = "RoundStarted"
	SubmitBid        MessageType = "SubmitBid"
	BidSubmitted     MessageType = "BidSubmitted"
	Tob              MessageType = "Tob"
)

type Message struct {
	Type    string `json:"type"`
	Payload []byte `json:"payload"`
}

type SendableMessage interface {
	MessageType() MessageType
	Marshal() ([]byte, error)
}

type ValidatableMessage interface {
	Validate() error
}

func validateRequiredFields(fields map[string]any) error {
	var missingFields []string
	for key, value := range fields {
		if value == nil {
			missingFields = append(missingFields, key)
		}
	}

	if len(missingFields) > 0 {
		return fmt.Errorf("missing required fields: %v", missingFields)
	}
	return nil
}

func ValidateMessage(msg ValidatableMessage) error {
	if msg == nil {
		return fmt.Errorf("message cannot be nil")
	}
	return msg.Validate()
}
