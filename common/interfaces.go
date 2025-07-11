package common

import "github.com/ethereum/go-ethereum/ethclient"

type Bidder interface {
	RpcNodeHttpClient() *ethclient.Client
	Nonce() uint64
	IncreaseNonce()
}
