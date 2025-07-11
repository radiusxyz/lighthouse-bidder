package manager

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/radiusxyz/lighthouse-bidder/config"
	"github.com/radiusxyz/lighthouse-bidder/lighthousewsclient"
	"github.com/radiusxyz/lighthouse-bidder/rpcnodewsclient"
)

type Manager struct {
	lighthouseWsClient *lighthousewsclient.LighthouseWsClient
	rpcNodeWsClient    *rpcnodewsclient.RpcNodeWsClient
	rpcNodeHttpClient  *ethclient.Client
	nonce              uint64
}

func New(conf *config.Config, bidderAddress string, bidderPrivateKey string, rollupIds []string) (*Manager, error) {
	rpcNodeHttpClient, err := ethclient.Dial(conf.RpcNodeHttpUrl)
	if err != nil {
		return nil, err
	}

	manager := &Manager{
		rpcNodeHttpClient: rpcNodeHttpClient,
		nonce:             0,
	}

	rpcNodeWsClient, err := rpcnodewsclient.New("cluster1-1", manager, conf.RpcNodeWsUrl, conf.AnvilUrl, rpcNodeHttpClient)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to rpc node: %w", err)
	}
	fmt.Println("Connected to the WebSocket rpc node!")

	lighthouseWsClient, err := lighthousewsclient.New(manager, conf.LighthouseUrl, conf.RpcNodeHttpUrl, bidderAddress, bidderPrivateKey, rollupIds)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to lighthouse: %w", err)
	}

	fmt.Println("Connected to the WebSocket lighthouse!")

	manager.rpcNodeWsClient = rpcNodeWsClient
	manager.lighthouseWsClient = lighthouseWsClient
	return manager, nil
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
