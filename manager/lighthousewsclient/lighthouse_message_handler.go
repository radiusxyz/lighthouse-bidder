package lighthousewsclient

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/radiusxyz/lighthouse-bidder/manager/lighthousewsclient/messages"
	"github.com/radiusxyz/lighthouse-bidder/types"
	"log"
)

type LighthouseMessageHandler struct {
	serverConn    *websocket.Conn
	bidderAddress string
}

func NewLighthouseMessageHandler(serverConn *websocket.Conn, bidderAddress string) *LighthouseMessageHandler {
	return &LighthouseMessageHandler{
		serverConn:    serverConn,
		bidderAddress: bidderAddress,
	}
}

func (l *LighthouseMessageHandler) handleBidderRegisteredMessage(message *messages.BidderRegisteredMessage) error {
	if err := messages.ValidateMessage(message); err != nil {
		log.Println("Validation failed:", err)
	}

	fmt.Println("입찰자로 정상 등록되었습니다:", message)
	return nil
}

func (l *LighthouseMessageHandler) handleAuctionCreatedMessage(message *messages.AuctionCreatedMessage) error {
	if err := messages.ValidateMessage(message); err != nil {
		log.Println("Validation failed:", err)
	}

	fmt.Println("auction이 성공적으로 생성되었습니다. auctionId: ", message.AuctionId)
	return nil
}

func (l *LighthouseMessageHandler) handleBidSubmittedMessage(message *messages.BidSubmittedMessage) error {
	if err := messages.ValidateMessage(message); err != nil {
		log.Println("Validation failed:", err)
	}

	fmt.Println("handleBidSubmittedMessage:", message)
	return nil
}

func (l *LighthouseMessageHandler) handleRoundStartedMessage(message *messages.RoundStartedMessage) error {
	fmt.Println("handleRoundStartedMessage:", message)

	if err := messages.ValidateMessage(message); err != nil {
		log.Println("Validation failed:", err)
	}

	transaction := &types.Transaction{
		Hash: "0xhash",
	}
	res := &messages.SubmitBidMessage{
		Bidder:       l.bidderAddress,
		AuctionId:    *message.AuctionId,
		Round:        *message.Round,
		GasPrice:     10,
		Transactions: []*types.Transaction{transaction},
	}
	return l.sendMessage(res)
}

func (l *LighthouseMessageHandler) handleTobMessage(message *messages.TobMessage) error {
	if err := messages.ValidateMessage(message); err != nil {
		log.Println("Validation failed:", err)
	}

	fmt.Println("handleTobMessage:", message)
	return nil
}

func (l *LighthouseMessageHandler) HandleMessage(message any) error {
	switch t := message.(type) {
	case *messages.BidderRegisteredMessage:
		return l.handleBidderRegisteredMessage(t)
	case *messages.RoundStartedMessage:
		return l.handleRoundStartedMessage(t)
	case *messages.BidSubmittedMessage:
		return l.handleBidSubmittedMessage(t)
	case *messages.TobMessage:
		return l.handleTobMessage(t)
	default:
		log.Println("Unknown message type:", t)
	}
	return nil
}

func (l *LighthouseMessageHandler) sendMessage(message messages.SendableMessage) error {
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

	return l.serverConn.WriteMessage(websocket.BinaryMessage, data)
}
