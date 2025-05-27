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
	currentAuctionTobTxs               map[string][]string         // rollupId -> rawTransactions
	currentAuctionRoundMyTxs           map[string]map[int][]string // rollupId -> roundIndex -> txs
	currentAuctionConfirmedTxScanIndex map[string]int              // rollupId -> index
	currentAuctionMyTxCount            map[string]uint64           //rollupId -> auction tx count
	lighthouseWsClient                 *lighthousewsclient.LighthouseWsClient
	rpcNodeWsClient                    *rpcnodewsclient.RpcNodeWsClient
	rpcNodeHttpClient                  *ethclient.Client
}

func New(conf *config.Config, bidderAddress string, bidderPrivateKey string, rollupIds []string) (*Manager, error) {
	rpcNodeHttpClient, err := ethclient.Dial(conf.RpcNodeHttpUrl)
	if err != nil {
		return nil, err
	}

	manager := &Manager{
		currentAuctionTobTxs:               make(map[string][]string),
		currentAuctionRoundMyTxs:           make(map[string]map[int][]string),
		currentAuctionConfirmedTxScanIndex: make(map[string]int),
		currentAuctionMyTxCount:            make(map[string]uint64),
		rpcNodeHttpClient:                  rpcNodeHttpClient,
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

func (m *Manager) CurrentAuctionRoundMyTxs(rollupId string, roundIndex int) ([]string, error) {
	rounds, ok := m.currentAuctionRoundMyTxs[rollupId]
	if !ok {
		return []string{}, nil
	}

	txs, ok := rounds[roundIndex]
	if !ok {
		return nil, fmt.Errorf("roundIndex %d not found for rollupId '%s'", roundIndex, rollupId)
	}

	return txs, nil
}

func (m *Manager) CurrentAuctionConfirmedTxScanIndex(rollupId string) int {
	return m.currentAuctionConfirmedTxScanIndex[rollupId]
}

func (m *Manager) CurrentAuctionMyTxCount(rollupId string) uint64 {
	return m.currentAuctionMyTxCount[rollupId]
}

func (m *Manager) IncreaseCurrentAuctionMyTxCount(rollupId string, addedTxCount uint64) uint64 {
	m.currentAuctionMyTxCount[rollupId] += addedTxCount
	return m.currentAuctionMyTxCount[rollupId]
}

func (m *Manager) ResetCurrentAuctionMyInfo(rollupId string) {
	m.currentAuctionConfirmedTxScanIndex[rollupId] = 0
	m.currentAuctionMyTxCount[rollupId] = 0
	m.currentAuctionRoundMyTxs[rollupId] = make(map[int][]string)
}

func (m *Manager) SetCurrentAuctionRoundMyInfo(rollupId string, roundIndex int, scanIndex int, addedTxs []string) {
	if _, ok := m.currentAuctionRoundMyTxs[rollupId]; !ok {
		m.currentAuctionRoundMyTxs[rollupId] = make(map[int][]string)
	}
	m.currentAuctionRoundMyTxs[rollupId][roundIndex] = addedTxs
	m.currentAuctionConfirmedTxScanIndex[rollupId] = scanIndex
}

func (m *Manager) RpcNodeHttpClient() *ethclient.Client {
	return m.rpcNodeHttpClient
}

func (m *Manager) SaveCurrentAuctionTobTxs(rollupId string, tobTxs []string) error {
	if m.currentAuctionTobTxs[rollupId] != nil {
		return fmt.Errorf("rollup '%s' is already in currentAuctionTobTxs", rollupId)
	}

	m.currentAuctionTobTxs[rollupId] = tobTxs
	return nil
}

func (m *Manager) CurrentAuctionTobTxs(rollupId string) ([]string, error) {
	tobTxs, ok := m.currentAuctionTobTxs[rollupId]
	if !ok {
		return nil, fmt.Errorf("rollupId '%s' not found in currentAuctionTobTxs", rollupId)
	}

	return tobTxs, nil
}
