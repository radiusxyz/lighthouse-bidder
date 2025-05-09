package rpcnodewsclient

import (
	"fmt"
	"github.com/gorilla/websocket"
)

type RpcNodeMessageHandler struct {
	serverConn *websocket.Conn
}

func NewRpcNodeMessageHandler(serverConn *websocket.Conn) *RpcNodeMessageHandler {
	return &RpcNodeMessageHandler{
		serverConn: serverConn,
	}
}

func (r *RpcNodeMessageHandler) HandleEnvelope(envelope []byte) error {
	fmt.Println("Received envelope: ", string(envelope))
	return nil
}
