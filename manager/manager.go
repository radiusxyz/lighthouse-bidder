package manager

import (
	"context"
	"fmt"
	"github.com/radiusxyz/lighthouse-bidder/anvilclient"
	"github.com/radiusxyz/lighthouse-bidder/config"
	"github.com/radiusxyz/lighthouse-bidder/lighthousewsclient"
	"github.com/radiusxyz/lighthouse-bidder/rpcnodewsclient"
)

type Manager struct {
	myAuctionRoundTxs                  map[string]map[int][]string // auctionId -> roundIndex -> txs
	currentAuctionConfirmedTxScanIndex map[string]int              // auctionId -> index
	myCurrentAuctionTxCount            map[string]uint64
	lighthouseWsClient                 *lighthousewsclient.LighthouseWsClient
	rpcNodeWsClient                    *rpcnodewsclient.RpcNodeWsClient
	anvil                              *anvilclient.Anvil
}

func New(conf *config.Config, bidderAddress string, bidderPrivateKey string, rollupIds []string) (*Manager, error) {
	anvil, err := anvilclient.New(conf.RpcNodeHttpUrl)
	if err != nil {
		return nil, err
	}

	rpcNodeWsClient, err := rpcnodewsclient.New(conf.RpcNodeWsUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to rpc node: %w", err)
	}
	fmt.Println("Connected to the WebSocket rpc node!")

	manager := &Manager{
		myAuctionRoundTxs:                  make(map[string]map[int][]string),
		currentAuctionConfirmedTxScanIndex: make(map[string]int),
		myCurrentAuctionTxCount:            make(map[string]uint64),
		rpcNodeWsClient:                    rpcNodeWsClient,
		anvil:                              anvil,
	}

	lighthouseWsClient, err := lighthousewsclient.New(manager, conf.LighthouseUrl, conf.RpcNodeHttpUrl, bidderAddress, bidderPrivateKey, rollupIds)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to lighthouse: %w", err)
	}
	fmt.Println("Connected to the WebSocket lighthouse!")

	manager.lighthouseWsClient = lighthouseWsClient

	return manager, nil
}

func (m *Manager) Start(ctx context.Context) {
	m.lighthouseWsClient.Start(ctx)
	m.rpcNodeWsClient.Start(ctx)
}

func (m *Manager) MyAuctionRoundTxs(auctionId string, roundIndex int) []string {
	return m.myAuctionRoundTxs[auctionId][roundIndex]
}

func (m *Manager) CurrentAuctionConfirmedTxScanIndex(auctionId string) int {
	return m.currentAuctionConfirmedTxScanIndex[auctionId]
}

func (m *Manager) MyCurrentAuctionTxCount(auctionId string) uint64 {
	return m.myCurrentAuctionTxCount[auctionId]
}

func (m *Manager) IncreaseMyCurrentAuctionTxCount(auctionId string, addedTxCount uint64) uint64 {
	m.myCurrentAuctionTxCount[auctionId] += addedTxCount
	return m.myCurrentAuctionTxCount[auctionId]
}

func (m *Manager) ResetMyCurrentAuctionInfo(auctionId string) {
	m.currentAuctionConfirmedTxScanIndex[auctionId] = 0
	m.myCurrentAuctionTxCount[auctionId] = 0
}

func (m *Manager) SetMyCurrentRoundInfo(auctionId string, roundIndex int, scanIndex int, addedTxs []string) {
	m.currentAuctionConfirmedTxScanIndex[auctionId] = scanIndex
	m.myAuctionRoundTxs[auctionId] = make(map[int][]string)
	m.myAuctionRoundTxs[auctionId][roundIndex] = addedTxs
}

//func (m *Manager) InitializeCurrentRound(auctionId string) {
//	m.myAuctionRoundTxs[auctionId] = make(map[int][]string)
//}
