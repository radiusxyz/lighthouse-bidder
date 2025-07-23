package requests

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type SubmitBidRequest struct {
	BidderAddress   common.Address `json:"bidderAddress"`
	AuctionId       string         `json:"auctionId"`
	BidAmount       *big.Int       `json:"bidAmount"`
	MetaTxNonce     *big.Int       `json:"metaTxNonce"`
	RawTransactions [][]byte       `json:"rawTransactions"`
	TxHashes        [][32]byte     `json:"txHashes"`
	Signature       []byte         `json:"signature"`
}

func (r *SubmitBidRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}
