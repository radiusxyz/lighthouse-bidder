package lighthousewsclient

import (
	"context"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/radiusxyz/lighthouse-bidder/manager/lighthousewsclient/requests"
	"io"
	"log"
)

type LighthouseWsClient struct {
	conn          *websocket.Conn
	rollupId      string
	lighthouseUrl string
	bidderAddress string
	leaveCh       chan struct{}
	envelopeCh    chan []byte
	handler       *LighthouseMessageHandler
}

func NewLighthouseWsClient(lighthouseUrl string, bidderAddress string, rollupId string) (*LighthouseWsClient, error) {
	conn, _, err := websocket.DefaultDialer.Dial(lighthouseUrl, nil)
	if err != nil {
		return nil, err
	}

	return &LighthouseWsClient{
		conn:          conn,
		rollupId:      rollupId,
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

	registerBidderMessage := &requests.RegisterBidderRequest{
		BidderAddress: l.bidderAddress,
		RollupId:      l.rollupId,
	}
	if err := l.handler.SendMessage(requests.RegisterBidder, registerBidderMessage); err != nil {
		log.Println("Write error:", err)
	}
}

func (l *LighthouseWsClient) ReadMessage() {
	defer func() {
		l.leaveCh <- struct{}{}
	}()

	for {
		_, envelope, err := l.conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
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
			log.Println("connection to the server has been lost")
		case envelope := <-l.envelopeCh:
			if err := l.handler.HandleEnvelope(envelope); err != nil {
				fmt.Println("exception filter: ", err.Error()) //Todo
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
