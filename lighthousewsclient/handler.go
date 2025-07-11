package lighthousewsclient

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	common2 "github.com/radiusxyz/lighthouse-bidder/common"
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
	bidder           common2.Bidder
}

func NewHandler(bidder common2.Bidder, serverConn *websocket.Conn, rpcNodeHttpUrl string, bidderAddress string, bidderPrivateKey string) (*LighthouseMessageHandler, error) {
	txBuilder, err := txbuilder.New(bidder.RpcNodeHttpClient(), rpcNodeHttpUrl)
	if err != nil {
		return nil, err
	}

	return &LighthouseMessageHandler{
		serverConn:       serverConn,
		bidderAddress:    bidderAddress,
		bidderPrivateKey: bidderPrivateKey,
		txBuilder:        txBuilder,
		bidder:           bidder,
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
	logger.Printf("Successfully bid submitted (auctionId: %s)", *resp.AuctionId)
	return nil
}

func (l *LighthouseMessageHandler) handleAuctionStartedEvent(event *events.AuctionStartedEvent) error {
	logger.ColorPrintf(logger.BgGreen, "Auction started (auctionId=%s)", *event.AuctionId)

	tx, err := l.txBuilder.GetSignedTransaction(l.bidderPrivateKey, common.HexToAddress("0x9965507D1a55bcC2695C58ba16FB37d819B0A4dc"), l.bidder.Nonce())
	if err != nil {
		return err
	}

	logger.Println("Transaction created")

	req := &requests.SubmitBidRequest{
		BidderAddress:   l.bidderAddress,
		AuctionId:       *event.AuctionId,
		BidPrice:        "1000000000000000000",
		RawTransactions: [][]byte{tx},
	}
	if err = l.SendMessage(requests.SubmitBid, req); err != nil {
		return err
	}

	l.bidder.IncreaseNonce()

	logger.Println("Bid submitted")
	return nil
}

func (l *LighthouseMessageHandler) handleTobEvent(event *events.TobEvent) error {
	logger.ColorPrintf(logger.BgGreen, "Received tob. rollupId '%s' auctionId '%s'", *event.RollupId, *event.AuctionId)
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
			return fmt.Errorf("failed to decode BidderVerifiedResponse: %w", err)
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
	case events.AuctionStarted:
		payload := new(events.AuctionStartedEvent)
		if err := json.Unmarshal(event.Payload, payload); err != nil {
			return fmt.Errorf("failed to decode AuctionStartedEvent: %w", err)
		}
		return l.handleAuctionStartedEvent(payload)
	case events.Tob:
		payload := new(events.TobEvent)
		if err := json.Unmarshal(event.Payload, payload); err != nil {
			return fmt.Errorf("failed to decode TobEvent: %w", err)
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
