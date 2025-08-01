package lighthousewsclient

import (
	"context"
	"encoding/base64"
	common2 "github.com/ethereum/go-ethereum/common"
	"github.com/gorilla/websocket"
	"github.com/radiusxyz/lighthouse-bidder/common"
	"github.com/radiusxyz/lighthouse-bidder/lighthousewsclient/requests"
	"github.com/radiusxyz/lighthouse-bidder/logger"
	"net/http"
	"strconv"
	"time"
)

type LighthouseWsClient struct {
	conn             *websocket.Conn
	rollupIds        []string
	lighthouseUrl    string
	bidderAddress    common2.Address
	bidderPrivateKey string
	leaveCh          chan struct{}
	envelopeCh       chan []byte
	handler          *LighthouseMessageHandler
}

func New(bidder common.Bidder, lighthouseUrl string, rpcNodeHttpUrl string, bidderAddress common2.Address, bidderPrivateKey string, rollupIds []string) (*LighthouseWsClient, error) {
	timestamp := uint64(time.Now().Unix())
	signature, err := common.GetSignature(bidderAddress.Hex(), timestamp, bidderPrivateKey)
	if err != nil {
		panic(err)
	}

	headers := http.Header{}
	headers.Set("Client-Type", "Bidder")
	headers.Set("Bidder-Address", bidderAddress.Hex())
	headers.Set("Signature", base64.StdEncoding.EncodeToString(signature))
	headers.Set("Timestamp", strconv.FormatUint(timestamp, 10))

	conn, _, err := websocket.DefaultDialer.Dial(lighthouseUrl, headers)
	if err != nil {
		return nil, err
	}

	handler, err := NewHandler(bidder, conn, rpcNodeHttpUrl, bidderAddress, bidderPrivateKey)
	if err != nil {
		return nil, err
	}

	return &LighthouseWsClient{
		conn:             conn,
		rollupIds:        rollupIds,
		bidderAddress:    bidderAddress,
		bidderPrivateKey: bidderPrivateKey,
		lighthouseUrl:    lighthouseUrl,
		leaveCh:          make(chan struct{}),
		envelopeCh:       make(chan []byte),
		handler:          handler,
	}, nil
}

func (l *LighthouseWsClient) Start(ctx context.Context) {
	for i := 0; i < 1; i++ {
		go l.ManageCh()
	}

	go l.ReadMessage()

	subscribeRollupsRequest := &requests.SubscribeRollupsRequest{
		BidderAddress: l.bidderAddress,
		RollupIds:     l.rollupIds,
	}

	if err := l.handler.SendMessage(requests.SubscribeRollups, subscribeRollupsRequest); err != nil {
		logger.ColorPrintf(logger.Red, "Error: %s\n", err.Error())
	}
}

func (l *LighthouseWsClient) ReadMessage() {
	for {
		_, envelope, err := l.conn.ReadMessage()
		if err != nil {
			logger.Println("Read error:", err)
			l.leaveCh <- struct{}{}
			break
		}
		l.envelopeCh <- envelope
	}
}

func (l *LighthouseWsClient) ManageCh() {
	for {
		select {
		case <-l.leaveCh:
			_ = l.conn.Close()
			logger.Println("connection to the server has been lost")
			l.Reconnect()

		case envelope := <-l.envelopeCh:
			if err := l.handler.HandleEnvelope(envelope); err != nil {
				logger.ColorPrintf(logger.Yellow, "Exception filter: %s\n", err.Error())
			}
		}
	}
}

func (l *LighthouseWsClient) Reconnect() {
	for {
		time.Sleep(time.Second * 5)
		timestamp := uint64(time.Now().Unix())
		signature, err := common.GetSignature(l.bidderAddress.Hex(), timestamp, l.bidderPrivateKey)
		if err != nil {
			panic(err)
		}

		headers := http.Header{}
		headers.Set("Client-Type", "Bidder")
		headers.Set("Bidder-Address", l.bidderAddress.Hex())
		headers.Set("Signature", base64.StdEncoding.EncodeToString(signature))
		headers.Set("Timestamp", strconv.FormatUint(timestamp, 10))

		conn, _, err := websocket.DefaultDialer.Dial(l.lighthouseUrl, headers)
		if err != nil {
			logger.ColorPrintf(logger.Red, "Dial error: %s", err.Error())
			continue
		}
		l.resetConn(conn)
		go l.ReadMessage()

		subscribeRollupsRequest := &requests.SubscribeRollupsRequest{
			BidderAddress: l.bidderAddress,
			RollupIds:     l.rollupIds,
		}

		if err := l.handler.SendMessage(requests.SubscribeRollups, subscribeRollupsRequest); err != nil {
			logger.ColorPrintf(logger.Red, "Error: %s\n", err.Error())
		}

		break
	}
}

func (l *LighthouseWsClient) resetConn(conn *websocket.Conn) {
	l.conn = conn
	l.handler.ResetConn(conn)
}

func (l *LighthouseWsClient) Close() error {
	return l.conn.Close()
}

func (l *LighthouseWsClient) WriteCloseMessage() error {
	return l.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
}
