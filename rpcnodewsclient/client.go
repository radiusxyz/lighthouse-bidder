package rpcnodewsclient

import (
	"context"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/radiusxyz/lighthouse-bidder/logger"
	"io"
)

type RpcNodeWsClient struct {
	conn       *websocket.Conn
	leaveCh    chan struct{}
	envelopeCh chan []byte
	handler    *RpcNodeMessageHandler
}

func New(rpcNodeUrl string) (*RpcNodeWsClient, error) {
	conn, _, err := websocket.DefaultDialer.Dial(rpcNodeUrl, nil)
	if err != nil {
		return nil, err
	}

	return &RpcNodeWsClient{
		conn:       conn,
		leaveCh:    make(chan struct{}),
		envelopeCh: make(chan []byte),
		handler:    NewRpcNodeMessageHandler(conn),
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
		"params": []interface{}{"newHeads"},
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
