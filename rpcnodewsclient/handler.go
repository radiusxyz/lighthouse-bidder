package rpcnodewsclient

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/goccy/go-json"
	"github.com/gorilla/websocket"
	"github.com/radiusxyz/lighthouse-bidder/common"
	"github.com/radiusxyz/lighthouse-bidder/logger"
	"github.com/radiusxyz/lighthouse-bidder/rpcnodewsclient/events"
	"math/big"
	"strings"
)

type WsMessage struct {
	JsonRpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  struct {
		Subscription string `json:"subscription"`
		Result       []byte `json:"result"`
	} `json:"params"`
}

type RpcNodeMessageHandler struct {
	rollupId          string
	serverConn        *websocket.Conn
	anvilClient       *ethclient.Client
	rpcNodeHttpClient *ethclient.Client
	lastBlockNumber   uint64
	bidder            common.Bidder
}

func NewRpcNodeMessageHandler(rollupId string, bidder common.Bidder, serverConn *websocket.Conn, rpcNodeHttpClient *ethclient.Client, anvilUrl string) (*RpcNodeMessageHandler, error) {
	anvilClient, err := ethclient.Dial(anvilUrl)
	if err != nil {
		return nil, err
	}
	return &RpcNodeMessageHandler{
		rollupId:          rollupId,
		serverConn:        serverConn,
		rpcNodeHttpClient: rpcNodeHttpClient,
		anvilClient:       anvilClient,
		bidder:            bidder,
	}, nil
}

func (r *RpcNodeMessageHandler) HandleEnvelope(envelope []byte) error {
	var msg WsMessage
	if err := json.Unmarshal(envelope, &msg); err != nil {
		return err
	}

	switch msg.Method {
	case "eth_subscription":
		var slotTransactions *events.SlotTransactions
		if err := json.Unmarshal(msg.Params.Result, &slotTransactions); err != nil {
			return err
		}

		logger.ColorPrintln(logger.Yellow, "SlotNumber: ", slotTransactions.SlotNumber)
		for i, tx := range slotTransactions.RawTransactions {
			logger.ColorPrintln(logger.Yellow, "tx ", i, ": ", tx)
		}

		// TODO: Apply on own tool and manage nonce
	}
	return nil
}

func (r *RpcNodeMessageHandler) IsOwnTx() bool {
	return true
}

func hexToUint64(hexStr string) (uint64, error) {
	n := new(big.Int)
	n.SetString(strings.TrimPrefix(hexStr, "0x"), 16)
	return n.Uint64(), nil
}
