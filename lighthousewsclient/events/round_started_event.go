package events

type RoundStartedEvent struct {
	RollupId              *string  `json:"rollupId"`
	AuctionId             *string  `json:"auctionId"`
	Round                 *int     `json:"round"`
	ConfirmedTransactions []string `json:"confirmedTransactions"`
}
