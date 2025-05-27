package events

type TobEvent struct {
	RollupId              *string  `json:"rollupId"`
	AuctionId             *string  `json:"auctionId"`
	SlotNumber            *int64   `json:"slotNumber"`
	ConfirmedTransactions []string `json:"confirmedTransactions"`
}
