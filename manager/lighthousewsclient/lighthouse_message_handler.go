package lighthousewsclient

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/radiusxyz/lighthouse-bidder/logger"
	"github.com/radiusxyz/lighthouse-bidder/manager/lighthousewsclient/events"
	"github.com/radiusxyz/lighthouse-bidder/manager/lighthousewsclient/requests"
	"github.com/radiusxyz/lighthouse-bidder/manager/lighthousewsclient/responses"
	"strconv"
)

type BaseMessage struct {
	Id        *string           `json:"id"`
	EventType *events.EventType `json:"eventType"`
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

func (l *LighthouseMessageHandler) handleBidderVerifiedResponse(res *responses.BidderVerifiedResponse) error {
	logger.ColorPrintln(logger.BgGreen, "Successfully verified")
	return nil
}

func (l *LighthouseMessageHandler) handleRollupsSubscribedResponse(res *responses.RollupsSubscribedResponse) error {
	logger.ColorPrintln(logger.BgGreen, "Successfully subscribed")
	return nil
}

func (l *LighthouseMessageHandler) handleRollupsUnsubscribedResponse(res *responses.RollupsUnsubscribedResponse) error {
	logger.ColorPrintln(logger.BgGreen, "Successfully unsubscribed")
	return nil
}

func (l *LighthouseMessageHandler) handleAllRollupsUnsubscribedResponse(res *responses.AllRollupsUnsubscribedResponse) error {
	logger.ColorPrintln(logger.BgGreen, "Successfully all unsubscribed")
	return nil
}

func (l *LighthouseMessageHandler) handleBidSubmittedResponse(res *responses.BidSubmittedResponse) error {
	logger.ColorPrintln(logger.BgGreen, "Successfully bid sent")
	return nil
}

func (l *LighthouseMessageHandler) handleRoundStartedEvent(event *events.RoundStartedEvent) error {
	logger.ColorPrintln(logger.BgGreen, "Round "+strconv.Itoa(*event.Round)+" started")

	transaction := "0xTOB" + *event.AuctionId + strconv.Itoa(*event.Round) + l.bidderAddress

	req := &requests.SubmitBidRequest{
		BidderAddress: l.bidderAddress,
		AuctionId:     *event.AuctionId,
		Round:         *event.Round,
		BidPrice:      10,
		Transactions:  []string{transaction},
	}
	if err := l.SendMessage(requests.SubmitBid, req); err != nil {
		return err
	}

	logger.Println("Bid submitted")

	return nil
}

func (l *LighthouseMessageHandler) handleTobEvent(event *events.TobEvent) error {
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
		res := new(responses.ResponseMessage)
		if err := json.Unmarshal(envelope, &res); err != nil {
			return fmt.Errorf("failed to parse response: %w", err)
		}
		if err := l.handleResponse(res); err != nil {
			return err
		}
	case base.EventType != nil:
		event := new(events.EventMessage)
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

func (l *LighthouseMessageHandler) handleResponse(res *responses.ResponseMessage) error {
	if res.Error != nil {
		return fmt.Errorf("[ErrorResponse] id=%s type=%s msg=%s", res.Id, res.ResponseType, res.Error.Message)
	}

	switch res.ResponseType {
	case responses.BidderVerified:
		payload := new(responses.BidderVerifiedResponse)
		if err := json.Unmarshal(res.Payload, payload); err != nil {
			return fmt.Errorf("failed to decode BidderRegisteredResponse: %w", err)
		}
		return l.handleBidderVerifiedResponse(payload)
	case responses.RollupsSubscribed:
		payload := new(responses.RollupsSubscribedResponse)
		if err := json.Unmarshal(res.Payload, payload); err != nil {
			return fmt.Errorf("failed to decode RollupsSubscribedResponse: %w", err)
		}
		return l.handleRollupsSubscribedResponse(payload)
	case responses.RollupsUnsubscribed:
		payload := new(responses.RollupsUnsubscribedResponse)
		if err := json.Unmarshal(res.Payload, payload); err != nil {
			return fmt.Errorf("failed to decode RollupsUnsubscribedResponse: %w", err)
		}
		return l.handleRollupsUnsubscribedResponse(payload)
	case responses.AllRollupsUnsubscribed:
		payload := new(responses.AllRollupsUnsubscribedResponse)
		if err := json.Unmarshal(res.Payload, payload); err != nil {
			return fmt.Errorf("failed to decode AllRollupsUnsubscribedResponse: %w", err)
		}
		return l.handleAllRollupsUnsubscribedResponse(payload)
	case responses.BidSubmitted:
		payload := new(responses.BidSubmittedResponse)
		if err := json.Unmarshal(res.Payload, payload); err != nil {
			return fmt.Errorf("failed to decode BidSubmittedResponse: %w", err)
		}
		return l.handleBidSubmittedResponse(payload)
	default:
		return fmt.Errorf("unknown response type")
	}
}

func (l *LighthouseMessageHandler) handleEvent(event *events.EventMessage) error {
	switch event.EventType {
	case events.RoundStarted:
		payload := new(events.RoundStartedEvent)
		if err := json.Unmarshal(event.Payload, payload); err != nil {
			return fmt.Errorf("failed to decode RoundStartedMessage: %w", err)
		}
		return l.handleRoundStartedEvent(payload)
	case events.Tob:
		payload := new(events.TobEvent)
		if err := json.Unmarshal(event.Payload, payload); err != nil {
			return fmt.Errorf("failed to decode BidSubmittedMessage: %w", err)
		}
		return l.handleTobEvent(payload)
	default:
		return fmt.Errorf("unknown event type")
	}
}

func (l *LighthouseMessageHandler) SendMessage(requestType requests.RequestType, message requests.RequestParams) error {
	payload, err := message.Marshal()
	if err != nil {
		return fmt.Errorf("failed to serialize message: %w", err)
	}

	requestMessage := &requests.RequestMessage{
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
