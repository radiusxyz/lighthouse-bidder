package lighthousewsclient

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	events2 "github.com/radiusxyz/lighthouse-bidder/lighthousewsclient/events"
	requests2 "github.com/radiusxyz/lighthouse-bidder/lighthousewsclient/requests"
	responses2 "github.com/radiusxyz/lighthouse-bidder/lighthousewsclient/responses"
	"github.com/radiusxyz/lighthouse-bidder/logger"
	"strconv"
)

type BaseMessage struct {
	Id        *string            `json:"id"`
	EventType *events2.EventType `json:"eventType"`
}

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

func (l *LighthouseMessageHandler) handleBidderVerifiedResponse(resp *responses2.BidderVerifiedResponse) error {
	logger.Println("Successfully verified")
	return nil
}

func (l *LighthouseMessageHandler) handleRollupsSubscribedResponse(resp *responses2.RollupsSubscribedResponse) error {
	logger.Println("Successfully subscribed")
	return nil
}

func (l *LighthouseMessageHandler) handleRollupsUnsubscribedResponse(resp *responses2.RollupsUnsubscribedResponse) error {
	logger.ColorPrintln(logger.BgGreen, "Successfully unsubscribed")
	return nil
}

func (l *LighthouseMessageHandler) handleAllRollupsUnsubscribedResponse(resp *responses2.AllRollupsUnsubscribedResponse) error {
	logger.Println("Successfully all unsubscribe")
	return nil
}

func (l *LighthouseMessageHandler) handleBidSubmittedResponse(resp *responses2.BidSubmittedResponse) error {
	logger.Printf("Successfully bid submitted (auctionId: %s, round:%d)", *resp.AuctionId, *resp.Round)
	return nil
}

func (l *LighthouseMessageHandler) handleRoundStartedEvent(event *events2.RoundStartedEvent) error {
	logger.ColorPrintf(logger.BgGreen, "Round started (auctionId=%s, round=%d)", *event.AuctionId, *event.Round)

	transaction := "0xTOB" + *event.AuctionId + strconv.Itoa(*event.Round) + l.bidderAddress

	req := &requests2.SubmitBidRequest{
		BidderAddress: l.bidderAddress,
		AuctionId:     *event.AuctionId,
		Round:         *event.Round,
		BidPrice:      "1000000000000000000",
		Transactions:  []string{transaction},
	}
	if err := l.SendMessage(requests2.SubmitBid, req); err != nil {
		return err
	}

	logger.Println("Bid submitted")

	return nil
}

func (l *LighthouseMessageHandler) handleTobEvent(event *events2.TobEvent) error {
	logger.ColorPrintln(logger.BgGreen, "Received tob. auctionId "+*event.AuctionId)
	return nil
}

func (l *LighthouseMessageHandler) HandleEnvelope(envelope []byte) error {
	base := new(BaseMessage)
	if err := json.Unmarshal(envelope, base); err != nil {
		return err
	}

	switch {
	case base.Id != nil:
		res := new(responses2.ResponseMessage)
		if err := json.Unmarshal(envelope, &res); err != nil {
			return fmt.Errorf("failed to parse response: %w", err)
		}
		if err := l.handleResponse(res); err != nil {
			return err
		}
	case base.EventType != nil:
		event := new(events2.EventMessage)
		if err := json.Unmarshal(envelope, &event); err != nil {
			return fmt.Errorf("failed to parse event: %w", err)
		}
		if err := l.handleEvent(event); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unknown message format")
	}
	return nil
}

func (l *LighthouseMessageHandler) handleResponse(res *responses2.ResponseMessage) error {
	if res.Error != nil {
		return fmt.Errorf("[ErrorResponse] id=%s type=%s msg=%s", res.Id, res.ResponseType, res.Error.Message)
	}

	switch res.ResponseType {
	case responses2.BidderVerified:
		payload := new(responses2.BidderVerifiedResponse)
		if err := json.Unmarshal(res.Payload, payload); err != nil {
			return fmt.Errorf("failed to decode BidderRegisteredResponse: %w", err)
		}
		return l.handleBidderVerifiedResponse(payload)
	case responses2.RollupsSubscribed:
		payload := new(responses2.RollupsSubscribedResponse)
		if err := json.Unmarshal(res.Payload, payload); err != nil {
			return fmt.Errorf("failed to decode RollupsSubscribedResponse: %w", err)
		}
		return l.handleRollupsSubscribedResponse(payload)
	case responses2.RollupsUnsubscribed:
		payload := new(responses2.RollupsUnsubscribedResponse)
		if err := json.Unmarshal(res.Payload, payload); err != nil {
			return fmt.Errorf("failed to decode RollupsUnsubscribedResponse: %w", err)
		}
		return l.handleRollupsUnsubscribedResponse(payload)
	case responses2.AllRollupsUnsubscribed:
		payload := new(responses2.AllRollupsUnsubscribedResponse)
		if err := json.Unmarshal(res.Payload, payload); err != nil {
			return fmt.Errorf("failed to decode AllRollupsUnsubscribedResponse: %w", err)
		}
		return l.handleAllRollupsUnsubscribedResponse(payload)
	case responses2.BidSubmitted:
		payload := new(responses2.BidSubmittedResponse)
		if err := json.Unmarshal(res.Payload, payload); err != nil {
			return fmt.Errorf("failed to decode BidSubmittedResponse: %w", err)
		}
		return l.handleBidSubmittedResponse(payload)
	default:
		return fmt.Errorf("unknown response type")
	}
}

func (l *LighthouseMessageHandler) handleEvent(event *events2.EventMessage) error {
	switch event.EventType {
	case events2.RoundStarted:
		payload := new(events2.RoundStartedEvent)
		if err := json.Unmarshal(event.Payload, payload); err != nil {
			return fmt.Errorf("failed to decode RoundStartedMessage: %w", err)
		}
		return l.handleRoundStartedEvent(payload)
	case events2.Tob:
		payload := new(events2.TobEvent)
		if err := json.Unmarshal(event.Payload, payload); err != nil {
			return fmt.Errorf("failed to decode BidSubmittedMessage: %w", err)
		}
		return l.handleTobEvent(payload)
	default:
		return fmt.Errorf("unknown event type")
	}
}

func (l *LighthouseMessageHandler) SendMessage(requestType requests2.RequestType, message requests2.RequestParams) error {
	payload, err := message.Marshal()
	if err != nil {
		return fmt.Errorf("failed to serialize message: %w", err)
	}

	requestMessage := &requests2.RequestMessage{
		Id:          uuid.New().String(),
		RequestType: requestType,
		Payload:     payload,
	}
	data, err := json.Marshal(requestMessage)
	if err != nil {
		return fmt.Errorf("failed to wrap message: %w", err)
	}

	return l.serverConn.WriteMessage(websocket.BinaryMessage, data)
}
