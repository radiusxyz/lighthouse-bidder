package main

type Status int

const (
	Success Status = iota
	Failure
)

type MessageType string

const (
	RegisterBidder   MessageType = "RegisterBidder"
	BidderRegistered MessageType = "BidderRegistered"
	SubmitBid        MessageType = "SubmitBid"
	BidSubmitted     MessageType = "BidSubmitted"
	RoundStarted     MessageType = "RoundStarted"
	Tob              MessageType = "Tob"
)

type Message struct {
	Type    string `json:"type"`
	Payload []byte `json:"payload"`
}

type RegisterBidderMessage struct {
	Bidder   string `json:"seller"`
	RollupId string `json:"rollupId"`
}

type BidderRegisteredMessage struct {
	Status Status `json:"status"`
}

type AuctionCreatedMessage struct {
	AuctionId string `json:"auctionId"`
	Status    Status `json:"status"`
}

type RoundStartedMessage struct {
	AuctionId string `json:"auctionId"`
	Round     int    `json:"round"`
}

type SubmitBidMessage struct {
	Bidder       string         `json:"bidder"`
	AuctionID    string         `json:"auctionId"`
	Round        int            `json:"round"`
	GasPrice     int            `json:"gasPrice"`
	Transactions []*Transaction `json:"transactions"`
}

type BidSubmittedMessage struct {
	Status Status `json:"status"`
}

type TobMessage struct {
	AuctionId    string         `json:"auctionId"`
	Transactions []*Transaction `json:"transactions"`
}

type Transaction struct {
	Hash string
}
