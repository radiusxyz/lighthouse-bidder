package lighthousewsclient

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/radiusxyz/lighthouse-bidder/logger"
	"github.com/radiusxyz/lighthouse-bidder/manager/lighthousewsclient/messages"
	"log"
	"strconv"
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

	logger.ColorLog(logger.BgGreen, "Successfully registered")

	return nil
}

func (l *LighthouseMessageHandler) handleAuctionCreatedMessage(message *messages.AuctionCreatedMessage) error {
	if err := messages.ValidateMessage(message); err != nil {
		log.Println("Validation failed:", err)
	}

	fmt.Println("Auction created. auctionId: ", message.AuctionId)
	return nil
}

func (l *LighthouseMessageHandler) handleBidSubmittedMessage(message *messages.BidSubmittedMessage) error {
	if err := messages.ValidateMessage(message); err != nil {
		log.Println("Validation failed:", err)
	}

	//fmt.Println("handleBidSubmittedMessage:", message)
	return nil
}

func (l *LighthouseMessageHandler) handleRoundStartedMessage(message *messages.RoundStartedMessage) error {
	logger.ColorLog(logger.Blue, "Round "+strconv.Itoa(*message.Round)+" started")

	if err := messages.ValidateMessage(message); err != nil {
		log.Println("Validation failed:", err)
	}

	transaction := "0xTOB" + *message.AuctionId + strconv.Itoa(*message.Round) + l.bidderAddress

	msg := &messages.SubmitBidMessage{
		BidderAddress: l.bidderAddress,
		AuctionId:     *message.AuctionId,
		Round:         *message.Round,
		BidPrice:      10,
		Transactions:  []string{transaction},
	}

	if err := l.sendMessage(msg); err != nil {
		return err
	}

	logger.ColorLog(logger.Green, "Bid sent. AuctionId: "+msg.AuctionId+" Round: "+strconv.Itoa(*message.Round))
	return nil
}

func (l *LighthouseMessageHandler) handleTobMessage(message *messages.TobMessage) error {
	if err := messages.ValidateMessage(message); err != nil {
		log.Println("Validation failed:", err)
	}

	logger.ColorLog(logger.Blue, "Received tob. auctionId "+*message.AuctionId)

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
