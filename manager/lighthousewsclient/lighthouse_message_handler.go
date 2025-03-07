package lighthousewsclient

import (
	"fmt"
	"github.com/radiusxyz/lighthouse-bidder/manager/lighthousewsclient/messages"
	"log"
)

type LighthouseMessageHandler struct {
}

func NewLighthouseMessageHandler() *LighthouseMessageHandler {
	return &LighthouseMessageHandler{}
}

func (l *LighthouseMessageHandler) handleRollupRegisteredMessage(message *messages.RollupRegisteredMessage) error {
	if err := messages.ValidateMessage(message); err != nil {
		log.Println("Validation failed:", err)
	}

	fmt.Println("판매자로 정상 등록되었습니다:", message)
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

func (l *LighthouseMessageHandler) HandleMessage(message any) error {
	switch t := message.(type) {
	case *messages.RollupRegisteredMessage:
		return l.handleRollupRegisteredMessage(t)
	case *messages.AuctionCreatedMessage:
		return l.handleAuctionCreatedMessage(t)
	case *messages.BidSubmittedMessage:
		return l.handleBidSubmittedMessage(t)
	default:
		log.Println("Unknown message type:", t)
	}
	return nil
}
