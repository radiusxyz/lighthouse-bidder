package rpcnodewsclient

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/websocket"
	common2 "github.com/radiusxyz/lighthouse-bidder/common"
	"github.com/radiusxyz/lighthouse-bidder/logger"
	"io"
)

type RpcNodeWsClient struct {
	conn       *websocket.Conn
	leaveCh    chan struct{}
	envelopeCh chan []byte
	handler    *RpcNodeMessageHandler
}

func New(rollupId string, bidder common2.Bidder, rpcNodeWsUrl string, anvilUrl string, rpcNodeHttpClient *ethclient.Client) (*RpcNodeWsClient, error) {
	conn, _, err := websocket.DefaultDialer.Dial(rpcNodeWsUrl, nil)
	if err != nil {
		return nil, err
	}

	handler, err := NewRpcNodeMessageHandler(rollupId, bidder, conn, rpcNodeHttpClient, anvilUrl)
	if err != nil {
		return nil, err
	}

	return &RpcNodeWsClient{
		conn:       conn,
		leaveCh:    make(chan struct{}),
		envelopeCh: make(chan []byte),
		handler:    handler,
	}, nil
}

func (r *RpcNodeWsClient) Start(ctx context.Context) {
	for i := 0; i < 1; i++ {
		go r.ManageCh()
	}

	go r.ReadMessage()

	subscribeReq := map[string]interface{}{
		"id":     1,
		"method": "eth_subscribe",
		"params": []interface{}{"newSlotTransactions"},
	}
	r.conn.WriteJSON(subscribeReq)
}

func (r *RpcNodeWsClient) ReadMessage() {
	defer func() {
		r.leaveCh <- struct{}{}
	}()

	for {
		_, envelope, err := r.conn.ReadMessage()
		if err != nil {
			logger.Println("Read error:", err)
			if errors.Is(err, io.EOF) {
				fmt.Println("Error eof")
				r.leaveCh <- struct{}{}
			}
			break
		}
		r.envelopeCh <- envelope
	}
}

func (r *RpcNodeWsClient) ManageCh() {
	for {
		select {
		case <-r.leaveCh:
			_ = r.conn.Close()
			logger.Println("connection to the server has been lost")
		case envelope := <-r.envelopeCh:
			if err := r.handler.HandleEnvelope(envelope); err != nil {
				logger.ColorPrintf(logger.Red, "Exception filter: %s\n", err.Error())
			}
		}
	}
}
