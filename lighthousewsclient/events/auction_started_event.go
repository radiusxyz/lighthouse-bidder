package events

type AuctionStartedEvent struct {
	RollupId   *string `json:"rollupId"`
	AuctionId  *string `json:"auctionId"`
	SlotNumber *int64  `json:"slotNumber"`
}
