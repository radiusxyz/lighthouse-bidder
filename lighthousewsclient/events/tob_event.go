package events

type TobEvent struct {
	AuctionId             *string  `json:"auctionId"`
	ConfirmedTransactions []string `json:"confirmedTransactions"`
}
