package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"time"
)

type Manager struct {
	Conn   *websocket.Conn
	Bidder string
}

func New() *Manager {
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
	if err != nil {
		log.Fatal("Failed to connect to lighthouse:", err)
	}
	fmt.Println("Connected to the WebSocket lighthouse!")

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	count := random.Intn(100)

	return &Manager{
		Conn:   conn,
		Bidder: "0xbidder" + strconv.Itoa(count),
	}
}

func (s *Manager) Start() {

	// 종료 시그널 처리
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// 메시지 수신 루프
	go s.ReadMessage()

	registerBidderMessage := &RegisterBidderMessage{
		Bidder:    s.Bidder,
		AuctionId: "https://youngmin.io_10",
	}
	if err := s.Write(registerBidderMessage); err != nil {
		log.Println("Write error:", err)
	}
	fmt.Println("register bidder completed")

	// 종료 처리
	select {
	case <-interrupt:
		fmt.Println("Interrupt received. Closing connection.")
		err := s.WriteCloseMessage()
		if err != nil {
			log.Println("Write close message error:", err)
		}
		if err = s.Close(); err != nil {
			log.Println("Close error:", err)
		}
		return
	}
}

func (s *Manager) decoder(message []byte) (*Message, error) {
	var msg = &Message{}
	if err := json.Unmarshal(message, msg); err != nil {
		log.Println("Invalid message format:", err)
		return nil, err
	}
	return msg, nil
}

func (s *Manager) Write(message any) error {
	data, err := MakeMessage(message)
	if err != nil {
		return err
	}
	if err = s.Conn.WriteMessage(websocket.BinaryMessage, data); err != nil {
		return err
	}
	return nil
}

func (s *Manager) ReadMessage() {
	for {
		_, message, err := s.Conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}
		msg, err := s.decoder(message)
		if err != nil {
			log.Println("Decode error:", err)
		}
		if err = s.routeMessage(msg); err != nil {
			log.Println("Route error:", err)
		}
	}
}

func (s *Manager) Close() error {
	return s.Conn.Close()
}

func (s *Manager) WriteCloseMessage() error {
	return s.Conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
}

func MakeAuctionId(rpcUrl string, chainId string) string {
	return rpcUrl + "_" + chainId
}

func MakeMessage(message any) ([]byte, error) {
	messageBytes, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}
	var messageType string
	switch message.(type) {
	case *RegisterBidderMessage:
		messageType = string(RegisterBidder)
	case *SubmitBidMessage:
		messageType = string(SubmitBid)
	}

	data := &Message{
		Type:    messageType,
		Payload: messageBytes,
	}
	ret, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (s *Manager) routeMessage(msg *Message) error {
	switch msg.Type {
	case "BidderRegistered":
		message := new(BidderRegisteredMessage)
		if err := UnMarshalFunc(msg.Payload, message); err != nil {
			return err
		}
		if err := s.handleBidderRegisteredMessage(message); err != nil {
			return err
		}
	case "ActionCreated":
		message := new(AuctionCreatedMessage)
		if err := UnMarshalFunc(msg.Payload, message); err != nil {
			return err
		}
		if err := s.handleAuctionCreatedMessage(message); err != nil {
			return err
		}
	case "BidSubmitted":
		message := new(BidSubmittedMessage)
		if err := UnMarshalFunc(msg.Payload, message); err != nil {
			return err
		}
		if err := s.handleBidSubmittedMessage(message); err != nil {
			return err
		}
	case "RoundStarted":
		message := new(RoundStartedMessage)
		if err := UnMarshalFunc(msg.Payload, message); err != nil {
			return err
		}
		if err := s.handleRoundStartedMessage(message); err != nil {
			return err
		}
	case "Tob":
		message := new(TobMessage)
		if err := UnMarshalFunc(msg.Payload, message); err != nil {
			return err
		}
		if err := s.handleTobMessage(message); err != nil {
			return err
		}
	default:
		log.Println("Unknown message type:", msg.Type)
	}
	return nil
}

func (s *Manager) handleBidderRegisteredMessage(message *BidderRegisteredMessage) error {
	fmt.Println("비더로 등록되었습니다:", message)
	return nil
}

func (s *Manager) handleAuctionCreatedMessage(message *AuctionCreatedMessage) error {
	fmt.Println("auction이 생성되었습니다. auctionId:", message.AuctionId)
	return nil
}

func (s *Manager) handleBidSubmittedMessage(message *BidSubmittedMessage) error {
	fmt.Println("비드가 제출되었습니다")
	return nil
}

func (s *Manager) handleRoundStartedMessage(message *RoundStartedMessage) error {
	fmt.Println("다음 라운드가 시작되었습니다. auctionId:", message.AuctionId, " round: ", message.Round)
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	gasPrice := random.Intn(100)
	submitBidMessage := &SubmitBidMessage{
		Bidder:       s.Bidder,
		AuctionID:    "https://youngmin.io_10",
		Round:        message.Round,
		GasPrice:     gasPrice,
		Transactions: s.getBundle(),
	}
	if err := s.Write(submitBidMessage); err != nil {
		log.Println("Write error:", err)
	}
	fmt.Println("비드를 제출하였습니다. auctionId:", message.AuctionId, " round:", message.Round)

	return nil
}

func (s *Manager) handleTobMessage(message *TobMessage) error {
	fmt.Println("Tob를 받았습니다.", message.AuctionId, " transactions: ", message.Transactions)
	return nil
}

func (s *Manager) getBundle() []*Transaction {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	count := random.Intn(3)
	txs := make([]*Transaction, count)
	for i := 0; i < count; i++ {
		txs[i] = &Transaction{
			Hash: "0x" + strconv.Itoa(count),
		}
	}
	return txs
}

func UnMarshalFunc[T any](payload []byte, result T) error {
	if err := json.Unmarshal(payload, result); err != nil {
		log.Println("Invalid message format:", err)
		return err
	}
	return nil
}
