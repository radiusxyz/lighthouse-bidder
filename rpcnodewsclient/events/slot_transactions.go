package events

type SlotTransactions struct {
	SlotNumber      int64    `json:"slotNumber"`
	RawTransactions []string `json:"rawTransactions"`
}
