package rpcnodewsclient

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/goccy/go-json"
	"github.com/gorilla/websocket"
	"github.com/radiusxyz/lighthouse-bidder/logger"
	"math/big"
	"strings"
)

type WsMessage struct {
	Method string `json:"method"`
	Params struct {
		Result struct {
			Number string `json:"number"`
		} `json:"result"`
	} `json:"params"`
}

type RpcNodeMessageHandler struct {
	serverConn        *websocket.Conn
	anvilClient       *ethclient.Client
	rpcNodeHttpClient *ethclient.Client
	lastBlockNumber   uint64
}

func NewRpcNodeMessageHandler(serverConn *websocket.Conn, rpcNodeHttpClient *ethclient.Client, anvilUrl string) (*RpcNodeMessageHandler, error) {
	anvilClient, err := ethclient.Dial(anvilUrl)
	if err != nil {
		return nil, err
	}
	return &RpcNodeMessageHandler{
		serverConn:        serverConn,
		rpcNodeHttpClient: rpcNodeHttpClient,
		anvilClient:       anvilClient,
	}, nil
}

func (r *RpcNodeMessageHandler) HandleEnvelope(envelope []byte) error {
	fmt.Println("Received envelope: ", string(envelope))
	var msg WsMessage
	if err := json.Unmarshal(envelope, &msg); err != nil {
		return err
	}

	switch msg.Method {
	case "eth_subscription":
		logger.ColorPrintln(logger.Yellow, "event msg: ", msg)
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
