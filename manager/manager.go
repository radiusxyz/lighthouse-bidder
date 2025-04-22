package manager

import (
	"context"
	"fmt"
	"github.com/radiusxyz/lighthouse-bidder/config"
	"github.com/radiusxyz/lighthouse-bidder/manager/lighthousewsclient"
)

type Manager struct {
	LighthouseWsClient *lighthousewsclient.LighthouseWsClient
}

func New(conf *config.Config, bidderAddress string, bidderPrivateKey string, rollupIds []string) (*Manager, error) {
	LighthouseWsClient, err := lighthousewsclient.NewLighthouseWsClient(conf.LighthouseUrl, bidderAddress, bidderPrivateKey, rollupIds)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to lighthouse: %w", err)
	}
	fmt.Println("Connected to the WebSocket lighthouse!")

	return &Manager{
		LighthouseWsClient: LighthouseWsClient,
	}, nil
}

func (m *Manager) Start(ctx context.Context) {
	m.LighthouseWsClient.Start(ctx)
}
