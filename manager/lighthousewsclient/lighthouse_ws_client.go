package lighthousewsclient

import (
	"context"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/radiusxyz/lighthouse-bidder/logger"
	"github.com/radiusxyz/lighthouse-bidder/manager/lighthousewsclient/requests"
	"io"
)

type LighthouseWsClient struct {
	conn          *websocket.Conn
	rollupIds     []string
	lighthouseUrl string
	bidderAddress string
	leaveCh       chan struct{}
	envelopeCh    chan []byte
	handler       *LighthouseMessageHandler
}

func NewLighthouseWsClient(lighthouseUrl string, bidderAddress string, rollupId []string) (*LighthouseWsClient, error) {
	conn, _, err := websocket.DefaultDialer.Dial(lighthouseUrl, nil)
	if err != nil {
		return nil, err
	}

	return &LighthouseWsClient{
		conn:          conn,
		rollupIds:     rollupId,
		bidderAddress: bidderAddress,
		lighthouseUrl: lighthouseUrl,
		leaveCh:       make(chan struct{}),
		envelopeCh:    make(chan []byte),
		handler:       NewLighthouseMessageHandler(conn, bidderAddress),
	}, nil
}

func (l *LighthouseWsClient) Start(ctx context.Context) {
	for i := 0; i < 1; i++ {
		go l.ManageCh()
	}

	go l.ReadMessage()

	verifyBidderRequest := &requests.VerifyBidderRequest{
		BidderAddress: l.bidderAddress,
	}
	if err := l.handler.SendMessage(requests.VerifyBidder, verifyBidderRequest); err != nil {
		logger.ColorPrintf(logger.Red, "Error: %s\n", err.Error())
	}

	subscribeRollupsRequest := &requests.SubscribeRollupsRequest{
		BidderAddress: l.bidderAddress,
		RollupIds:     l.rollupIds,
	}
	if err := l.handler.SendMessage(requests.SubscribeRollups, subscribeRollupsRequest); err != nil {
		logger.ColorPrintf(logger.Red, "Error: %s\n", err.Error())
	}
}

func (l *LighthouseWsClient) ReadMessage() {
	defer func() {
		l.leaveCh <- struct{}{}
	}()

	for {
		_, envelope, err := l.conn.ReadMessage()
		if err != nil {
			logger.Println("Read error:", err)
			if errors.Is(err, io.EOF) {
				fmt.Println("youngmin - eof")
				l.leaveCh <- struct{}{}
			}
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
		case envelope := <-l.envelopeCh:
			if err := l.handler.HandleEnvelope(envelope); err != nil {
				logger.ColorPrintf(logger.Red, "Exception filter: %s\n", err.Error())
			}
		}
	}
}

func (l *LighthouseWsClient) Close() error {
	return l.conn.Close()
}

func (l *LighthouseWsClient) WriteCloseMessage() error {
	return l.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
}
