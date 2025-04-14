package events

type RoundStartedEvent struct {
	AuctionId             *string  `json:"auctionId"`
	Round                 *int     `json:"round"`
	ConfirmedTransactions []string `json:"confirmedTransactions"`
}
