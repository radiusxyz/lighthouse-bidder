package manager

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/radiusxyz/lighthouse-bidder/config"
	"github.com/radiusxyz/lighthouse-bidder/lighthousewsclient"
	"github.com/radiusxyz/lighthouse-bidder/logger"
	"github.com/radiusxyz/lighthouse-bidder/rpcnodewsclient"
	"log"
	"math/big"
	"sync"
	"time"
)

type Manager struct {
	lighthouseWsClient    *lighthousewsclient.LighthouseWsClient
	rpcNodeWsClient       *rpcnodewsclient.RpcNodeWsClient
	rpcNodeHttpClient     *ethclient.Client
	nonce                 uint64
	metaTxNonce           *big.Int
	conf                  *config.Config
	bidderAddress         common.Address
	isMevCatchingMutex    sync.RWMutex
	auctionContractClient *ContractClient
}

func New(conf *config.Config, bidderAddress common.Address, bidderPrivateKey string, rollupIds []string) (*Manager, error) {
	rpcNodeHttpClient, err := ethclient.Dial(*conf.RpcNodeHttpUrl)
	if err != nil {
		return nil, err
	}

	nonce, err := rpcNodeHttpClient.PendingNonceAt(context.Background(), bidderAddress)
	if err != nil {
		log.Fatalf("failed to get nonce: %v", err)
	}
	logger.ColorPrintf(logger.BrightGreen, "NONCE: %d", nonce)

	contractClient, err := NewContractClient(conf)
	if err != nil {
		panic("failed to create contract client" + err.Error())
	}

	metaTxNonce, err := contractClient.GetNonce(bidderAddress)
	if err != nil {
		panic("failed to get nonce: " + err.Error())
	}
	logger.ColorPrintf(logger.BgYellow, "nonono: (%d)", metaTxNonce)
	manager := &Manager{
		rpcNodeHttpClient:     rpcNodeHttpClient,
		nonce:                 nonce,
		metaTxNonce:           metaTxNonce,
		conf:                  conf,
		bidderAddress:         bidderAddress,
		auctionContractClient: contractClient,
	}

	rpcNodeWsClient, err := rpcnodewsclient.New(*conf.RollupId, manager, *conf.RpcNodeWsUrl, *conf.AnvilUrl, rpcNodeHttpClient)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to rpc node: %w", err)
	}
	fmt.Println("Connected to the WebSocket rpc node!")

	lighthouseWsClient, err := lighthousewsclient.New(manager, *conf.LighthouseUrl, *conf.RpcNodeHttpUrl, bidderAddress, bidderPrivateKey, rollupIds)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to lighthouse: %w", err)
	}

	fmt.Println("Connected to the WebSocket lighthouse!")

	manager.rpcNodeWsClient = rpcNodeWsClient
	manager.lighthouseWsClient = lighthouseWsClient
	return manager, nil
}

func (m *Manager) Config() *config.Config {
	return m.conf
}

func (m *Manager) Start(ctx context.Context) {
	m.lighthouseWsClient.Start(ctx)
	m.rpcNodeWsClient.Start(ctx)
}

func (m *Manager) RpcNodeHttpClient() *ethclient.Client {
	return m.rpcNodeHttpClient
}

func (m *Manager) Nonce() uint64 {
	return m.nonce
}

func (m *Manager) IncreaseNonce() {
	m.nonce++
}

func (m *Manager) MetaTxNonce() *big.Int {
	return m.metaTxNonce
}

func (m *Manager) UpdateMetaTxNonce(succeed bool) {
	if succeed {
		m.metaTxNonce.Add(m.metaTxNonce, big.NewInt(1))
	} else {
		m.metaTxNonce = m.MetaTxNonce2()
	}
}

func (m *Manager) PendingNonceAt() uint64 {
	nonce, err := m.rpcNodeHttpClient.PendingNonceAt(context.Background(), m.bidderAddress)
	if err != nil {
		log.Fatalf("failed to get nonce: %v", err)
	}
	return nonce
}

func (m *Manager) SearchMev() {
	m.isMevCatchingMutex.Lock()
	defer m.isMevCatchingMutex.Unlock()

	time.Sleep(200 * time.Millisecond)
	logger.ColorPrintln(logger.BrightYellow, "Catch the MEV case")
}

func (m *Manager) WaitMevCatching() {
	m.isMevCatchingMutex.RLock()
	defer m.isMevCatchingMutex.RUnlock()
}

func (m *Manager) MetaTxNonce2() *big.Int {
	nonce, err := m.auctionContractClient.GetNonce(m.bidderAddress)
	if err != nil {
		log.Fatalf("failed to get nonce: %v", err)
	}
	return nonce
}
