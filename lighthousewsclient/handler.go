package lighthousewsclient

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/radiusxyz/lighthouse-bidder/lighthousewsclient/events"
	"github.com/radiusxyz/lighthouse-bidder/lighthousewsclient/requests"
	"github.com/radiusxyz/lighthouse-bidder/lighthousewsclient/responses"
	"github.com/radiusxyz/lighthouse-bidder/logger"
	"github.com/radiusxyz/lighthouse-bidder/txbuilder"
)

type BaseMessage struct {
	Id        *string           `json:"id"`
	EventType *events.EventType `json:"eventType"`
}

type LighthouseMessageHandler struct {
	serverConn       *websocket.Conn
	bidderAddress    string
	bidderPrivateKey string
	txBuilder        *txbuilder.TxBuilder
}

func NewHandler(serverConn *websocket.Conn, rpcNodeHttpUrl string, bidderAddress string, bidderPrivateKey string) (*LighthouseMessageHandler, error) {
	txBuilder, err := txbuilder.New(rpcNodeHttpUrl)
	if err != nil {
		return nil, err
	}

	return &LighthouseMessageHandler{
		serverConn:       serverConn,
		bidderAddress:    bidderAddress,
		bidderPrivateKey: bidderPrivateKey,
		txBuilder:        txBuilder,
	}, nil
}

func (l *LighthouseMessageHandler) handleBidderVerifiedResponse(resp *responses.BidderVerifiedResponse) error {
	logger.Println("Successfully verified")
	return nil
}

func (l *LighthouseMessageHandler) handleRollupsSubscribedResponse(resp *responses.RollupsSubscribedResponse) error {
	logger.Println("Successfully subscribed")
	return nil
}

func (l *LighthouseMessageHandler) handleRollupsUnsubscribedResponse(resp *responses.RollupsUnsubscribedResponse) error {
	logger.ColorPrintln(logger.BgGreen, "Successfully unsubscribed")
	return nil
}

func (l *LighthouseMessageHandler) handleAllRollupsUnsubscribedResponse(resp *responses.AllRollupsUnsubscribedResponse) error {
	logger.Println("Successfully all unsubscribe")
	return nil
}

func (l *LighthouseMessageHandler) handleBidSubmittedResponse(resp *responses.BidSubmittedResponse) error {
	logger.Printf("Successfully bid submitted (auctionId: %s, round:%d)", *resp.AuctionId, *resp.Round)
	return nil
}

func (l *LighthouseMessageHandler) handleRoundStartedEvent(event *events.RoundStartedEvent) error {
	logger.ColorPrintf(logger.BgGreen, "Round started (auctionId=%s, round=%d)", *event.AuctionId, *event.Round)

	hexTx, err := l.txBuilder.GetSignedTransaction(l.bidderPrivateKey, common.HexToAddress("0x9965507D1a55bcC2695C58ba16FB37d819B0A4dc"))
	if err != nil {
		return err
	}

	logger.Println("Transaction created")

	req := &requests.SubmitBidRequest{
		BidderAddress: l.bidderAddress,
		AuctionId:     *event.AuctionId,
		Round:         *event.Round,
		BidPrice:      "1000000000000000000",
		Transactions:  []string{hexTx},
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
