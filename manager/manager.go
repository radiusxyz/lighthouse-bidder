package manager

import (
	"context"
	"fmt"
	"github.com/radiusxyz/lighthouse-bidder/config"
	"github.com/radiusxyz/lighthouse-bidder/lighthousewsclient"
	"github.com/radiusxyz/lighthouse-bidder/rpcnodewsclient"
)

type Manager struct {
	lighthouseWsClient *lighthousewsclient.LighthouseWsClient
	rpcNodeWsClient    *rpcnodewsclient.RpcNodeWsClient
}

func New(conf *config.Config, bidderAddress string, bidderPrivateKey string, rollupIds []string) (*Manager, error) {
	rpcNodeWsClient, err := rpcnodewsclient.New(conf.RpcNodeWsUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to rpc node: %w", err)
	}
	fmt.Println("Connected to the WebSocket rpc node!")

	lighthouseWsClient, err := lighthousewsclient.New(conf.LighthouseUrl, conf.RpcNodeHttpUrl, bidderAddress, bidderPrivateKey, rollupIds)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to lighthouse: %w", err)
	}
	fmt.Println("Connected to the WebSocket lighthouse!")

	return &Manager{
		lighthouseWsClient: lighthouseWsClient,
		rpcNodeWsClient:    rpcNodeWsClient,
	}, nil
}

func (m *Manager) Start(ctx context.Context) {
	m.lighthouseWsClient.Start(ctx)
	m.rpcNodeWsClient.Start(ctx)
}
