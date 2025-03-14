package lighthousewsclient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/radiusxyz/lighthouse-bidder/manager/lighthousewsclient/messages"
	"io"
	"log"
)

type LighthouseWsClient struct {
	conn           *websocket.Conn
	rollupId       string
	lighthouseUrl  string
	bidderAddress  string
	leaveCh        chan struct{}
	envelopeCh     chan []byte
	messageDecoder EnvelopeDecodeFunc
	handler        *LighthouseMessageHandler
}

func NewLighthouseWsClient(lighthouseUrl string, bidderAddress string, rollupId string) (*LighthouseWsClient, error) {
	conn, _, err := websocket.DefaultDialer.Dial(lighthouseUrl, nil)
	if err != nil {
		return nil, err
	}

	return &LighthouseWsClient{
		conn:           conn,
		rollupId:       rollupId,
		bidderAddress:  bidderAddress,
		lighthouseUrl:  lighthouseUrl,
		leaveCh:        make(chan struct{}),
		envelopeCh:     make(chan []byte),
		messageDecoder: DecodeEnvelopeFunc,
		handler:        NewLighthouseMessageHandler(conn, bidderAddress),
	}, nil
}

func (l *LighthouseWsClient) Start(ctx context.Context) {
	for i := 0; i < 1; i++ {
		go l.ManageCh()
	}

	go l.ReadMessage()

	registerBidderMessage := &messages.RegisterBidderMessage{
		BidderAddress: l.bidderAddress,
		RollupId:      l.rollupId,
	}
	if err := l.SendMessage(registerBidderMessage); err != nil {
		log.Println("Write error:", err)
	}
}

func (l *LighthouseWsClient) SendMessage(message messages.SendableMessage) error {
	payload, err := message.Marshal()
	if err != nil {
		return fmt.Errorf("failed to serialize message: %w", err)
	}
	wrapper := &messages.Message{
		Type:    string(message.MessageType()),
		Payload: payload,
	}
	data, err := json.Marshal(wrapper)
	if err != nil {
		return fmt.Errorf("failed to wrap message: %w", err)
	}
	return l.conn.WriteMessage(websocket.BinaryMessage, data)
}

func (l *LighthouseWsClient) ReadMessage() {
	defer func() {
		l.leaveCh <- struct{}{}
	}()

	for {
		_, message, err := l.conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			if errors.Is(err, io.EOF) {
				fmt.Println("youngmin - eof")
				l.leaveCh <- struct{}{}
			}
			break
		}
		l.envelopeCh <- message
	}
}

func (l *LighthouseWsClient) ManageCh() {
	for {
		select {
		case <-l.leaveCh:
			_ = l.conn.Close()
			log.Println("connection to the server has been lost")
		case envelope := <-l.envelopeCh:
			decodedMessage, err := l.messageDecoder(envelope)
			if err != nil {
				fmt.Println(err) //Todo
			}
			if err = l.handler.HandleMessage(decodedMessage); err != nil {
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

type EnvelopeDecodeFunc func([]byte) (any, error)

func DecodeEnvelopeFunc(envelope []byte) (any, error) {
	message, err := decodePayload(envelope)
	if err != nil {
		return nil, err
	}

	var data any
	switch message.Type {
	case string(messages.BidderRegistered):
		data = new(messages.BidderRegisteredMessage)
		if err := UnMarshalFunc(message.Payload, data); err != nil {
			return nil, err
		}
	case string(messages.RoundStarted):
		data = new(messages.RoundStartedMessage)
		if err = UnMarshalFunc(message.Payload, data); err != nil {
			return nil, err
		}
	case string(messages.BidSubmitted):
		data = new(messages.BidSubmittedMessage)
		if err = UnMarshalFunc(message.Payload, data); err != nil {
			return nil, err
		}
	case string(messages.Tob):
		data = new(messages.TobMessage)
		if err = UnMarshalFunc(message.Payload, data); err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("invalid message header %x", message.Type)
	}

	return data, nil
}

func UnMarshalFunc[T any](payload []byte, result T) error {
	if err := json.Unmarshal(payload, result); err != nil {
		log.Println("Invalid message format:", err)
		return err
	}
	return nil
}

func decodePayload(message []byte) (*messages.Message, error) {
	var msg = &messages.Message{}
	if err := json.Unmarshal(message, msg); err != nil {
		log.Println("Invalid message format:", err)
		return nil, err
	}
	return msg, nil
}
