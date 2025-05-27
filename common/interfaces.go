package common

import "github.com/ethereum/go-ethereum/ethclient"

type Bidder interface {
	CurrentAuctionRoundMyTxs(rollupId string, roundIndex int) ([]string, error)
	CurrentAuctionConfirmedTxScanIndex(rollupId string) int
	CurrentAuctionMyTxCount(rollupId string) uint64
	IncreaseCurrentAuctionMyTxCount(rollupId string, addedTxCount uint64) uint64
	ResetCurrentAuctionMyInfo(rollupId string)
	SetCurrentAuctionRoundMyInfo(rollupId string, roundIndex int, scanIndex int, addedTxs []string)
	RpcNodeHttpClient() *ethclient.Client
	SaveCurrentAuctionTobTxs(rollupId string, tobTxs []string) error
	CurrentAuctionTobTxs(rollupId string) ([]string, error)
}
