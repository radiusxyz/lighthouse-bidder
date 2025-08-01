package common

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/radiusxyz/lighthouse-bidder/config"
	"math/big"
)

type Bidder interface {
	RpcNodeHttpClient() *ethclient.Client
	Nonce() uint64
	AuctionNonce() *big.Int
	IncreaseNonce()
	Config() *config.Config
	PendingNonceAt() uint64
	SearchMev()
	WaitMevCatching()
	FetchAuctionNonce() *big.Int
	UpdateAuctionNonce(flag bool)
}
