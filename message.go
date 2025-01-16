package main

type Status int

const (
	Success Status = iota
	Failure
)

type MessageType string

const (
	RegisterSeller   MessageType = "RegisterSeller"
	RegisterBidder   MessageType = "RegisterBidder"
	SellerRegistered MessageType = "SellerRegistered"
	CreateAuction    MessageType = "CreateAuction"
	AuctionCreated   MessageType = "AuctionCreated"
	SubmitBid        MessageType = "SubmitBid"
	BidSubmitted     MessageType = "BidSubmitted"
)

type Message struct {
	Type    string `json:"type"`
	Payload []byte `json:"payload"`
}

type RegisterSellerMessage struct {
	Seller string `json:"seller"`
}

type SellerRegisteredMessage struct {
	Status Status `json:"status"`
}

type RegisterBidderMessage struct {
	Bidder    string `json:"seller"`
	AuctionId string `json:"auctionId"`
}

type BidderRegisteredMessage struct {
	Status Status `json:"status"`
}

type CreateAuctionMessage struct {
	Seller      string `json:"seller"`
	RpcUrl      string `json:"rpc_url"`
	ChainID     string `json:"chain_id"`
	BlockNumber int64  `json:"block_number"`
	BlockTime   int64  `json:"block_time"`
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
	GasPrice     int            `json:"gas_price"`
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
