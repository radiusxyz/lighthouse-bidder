package rpcnodewsclient

import "github.com/gorilla/websocket"

type RpcNodeMessageHandler struct {
	serverConn *websocket.Conn
}

func NewRpcNodeMessageHandler(serverConn *websocket.Conn) *RpcNodeMessageHandler {
	return &RpcNodeMessageHandler{
		serverConn: serverConn,
	}
}

func (r *RpcNodeMessageHandler) HandleEnvelope(envelope []byte) error {
	return nil
}
