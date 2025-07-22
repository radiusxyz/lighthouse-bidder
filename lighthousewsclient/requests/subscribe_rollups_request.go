package requests

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
)

type SubscribeRollupsRequest struct {
	BidderAddress common.Address `json:"bidderAddress"`
	RollupIds     []string       `json:"rollupIds"`
}

func (r *SubscribeRollupsRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}
