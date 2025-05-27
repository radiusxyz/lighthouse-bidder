package rpcnodewsclient

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/goccy/go-json"
	"github.com/gorilla/websocket"
	common2 "github.com/radiusxyz/lighthouse-bidder/common"
	"github.com/radiusxyz/lighthouse-bidder/logger"
	"github.com/radiusxyz/lighthouse-bidder/rpcnodewsclient/events"
	"math/big"
	"strconv"
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
	bidder            common2.Bidder
}

func NewRpcNodeMessageHandler(rollupId string, bidder common2.Bidder, serverConn *websocket.Conn, rpcNodeHttpClient *ethclient.Client, anvilUrl string) (*RpcNodeMessageHandler, error) {
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

		tobTxs, err := r.bidder.CurrentAuctionTobTxs(r.rollupId + "_" + strconv.FormatInt(slotTransactions.SlotNumber, 10))
		if err != nil {
			return err
		}

		bobTxs := slotTransactions.RawTransactions[len(tobTxs):]
		for i, tx := range bobTxs {
			logger.ColorPrintln(logger.Yellow, "bob ", i, ": ", tx)
		}
		//r.bidder.RollupSlotTobTxs()
		//receivedBlockNumber, err := hexToUint64(msg.Params.Result.Number)
		//if err != nil {
		//	return err
		//}

		//if r.lastBlockNumber+1 != receivedBlockNumber {
		//	panic("can not process block")
		//}
		//fmt.Println("aaaa: ", msg.Params.Result.Number)
		//blockNumber := new(big.Int)
		//blockNumber.SetString(strings.TrimPrefix(msg.Params.Result.Number, "0x"), 16)
		//block, err := r.rpcNodeHttpClient.BlockByNumber(context.Background(), blockNumber)
		//if err != nil {
		//	return err
		//}
		//for _, tx := range block.Transactions() {
		//	fmt.Println("ym: ", tx)
		//}
	}
	return nil
}

func hexToUint64(hexStr string) (uint64, error) {
	n := new(big.Int)
	n.SetString(strings.TrimPrefix(hexStr, "0x"), 16)
	return n.Uint64(), nil
}
