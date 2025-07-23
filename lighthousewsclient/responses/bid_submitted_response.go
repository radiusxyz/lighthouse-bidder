package responses

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type BidSubmittedResponse struct {
	BidderAddress   *common.Address `json:"bidderAddress"`
	AuctionId       *string         `json:"auctionId"`
	BidAmount       *big.Int        `json:"bidAmount"`
	MetaTxNonce     *big.Int        `json:"metaTxNonce"`
	RawTransactions [][]byte        `json:"rawTransactions"`
	TxHashes        [][32]byte      `json:"txHashes"`
	Signature       []byte          `json:"signature"`
}

func (r *BidSubmittedResponse) Unmarshal(data []byte) error {
	if err := json.Unmarshal(data, r); err != nil {
		return err
	}
	return nil
}
