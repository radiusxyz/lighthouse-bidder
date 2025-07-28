package lighthousewsclient

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	common2 "github.com/radiusxyz/lighthouse-bidder/common"
	"github.com/radiusxyz/lighthouse-bidder/lighthousewsclient/events"
	"github.com/radiusxyz/lighthouse-bidder/lighthousewsclient/requests"
	"github.com/radiusxyz/lighthouse-bidder/lighthousewsclient/responses"
	"github.com/radiusxyz/lighthouse-bidder/logger"
	"github.com/radiusxyz/lighthouse-bidder/txbuilder"
	"log"
	"math/big"
	"math/rand"
	"strings"
	"time"
)

type BaseMessage struct {
	Id        *string           `json:"id"`
	EventType *events.EventType `json:"eventType"`
}

type LighthouseMessageHandler struct {
	serverConn       *websocket.Conn
	bidderAddress    common.Address
	bidderPrivateKey *ecdsa.PrivateKey
	txBuilder        *txbuilder.TxBuilder
	bidder           common2.Bidder
}

func NewHandler(bidder common2.Bidder, serverConn *websocket.Conn, rpcNodeHttpUrl string, bidderAddress common.Address, bidderPrivateKey string) (*LighthouseMessageHandler, error) {
	txBuilder, err := txbuilder.New(bidder.RpcNodeHttpClient(), rpcNodeHttpUrl)
	if err != nil {
		return nil, err
	}

	privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(bidderPrivateKey, "0x"))
	if err != nil {
		log.Fatalf("Failed to load private key: %v", err)

	}

	return &LighthouseMessageHandler{
		serverConn:       serverConn,
		bidderAddress:    bidderAddress,
		bidderPrivateKey: privateKey,
		txBuilder:        txBuilder,
		bidder:           bidder,
	}, nil
}

func (l *LighthouseMessageHandler) ResetConn(conn *websocket.Conn) {
	l.serverConn = conn
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

	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	l.bidder.IsMevCatching()

	signedTx, err := l.txBuilder.GetSignedTransaction(l.bidderPrivateKey, address, l.bidder.PendingNonceAt())
	if err != nil {
		return err
	}

	rawTx, err := signedTx.MarshalBinary()
	if err != nil {
		log.Fatalf("failed to encode tx: %v", err)
	}

	logger.Println("Transaction created")

	txHashes := ConvertToBytes32Array([]common.Hash{
		signedTx.Hash(),
	})

	rawTypeHash := crypto.Keccak256([]byte("SubmitBid(uint256 bidPrice,uint256 nonce,bytes32 bidTxdata)"))
	nonce := big.NewInt(1)

	var typeHashForpacking [32]byte
	if len(rawTypeHash) != 32 {
		log.Fatalf("Invalid type hash length: %d", len(rawTypeHash))
	}
	copy(typeHashForpacking[:], rawTypeHash)

	var txHashesArray []common.Hash
	for _, txHash := range txHashesArray {

		txHashesArray = append(txHashesArray, txHash)
	}

	packedBytes := make([]byte, 0, len(txHashesArray)*common.HashLength)
	for _, txHash := range txHashesArray {
		packedBytes = append(packedBytes, txHash.Bytes()...)
	}

	finalBidTxdata := crypto.Keccak256Hash(packedBytes)

	bytes32Ty, err := abi.NewType("bytes32", "", nil)
	if err != nil {
		log.Fatalf("Failed to create bytes32 type: %v", err)

	}
	uint256Ty, err := abi.NewType("uint256", "", nil)
	if err != nil {
		log.Fatalf("Failed to create uint256 type: %v", err)
	}
	addressTy, err := abi.NewType("address", "", nil)
	if err != nil {
		log.Fatalf("Failed to create address type: %v", err)
	}

	arguments := abi.Arguments{
		{Type: bytes32Ty}, // typeHash
		{Type: uint256Ty}, // price1
		{Type: uint256Ty}, // nonce
		{Type: bytes32Ty}, // bidTxdata

	}

	var bidTxdataForPacking [32]byte = finalBidTxdata

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	bidAmount := big.NewInt(int64(10000000000 + r.Intn(100)))

	packed, err := arguments.Pack(
		typeHashForpacking,
		bidAmount,
		nonce,
		bidTxdataForPacking,
	)

	if err != nil {
		log.Fatalf("Failed to pack arguments: %v", err)
	}

	structHash := crypto.Keccak256Hash(packed)

	domainTypeHash := crypto.Keccak256Hash([]byte("EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)"))
	nameHash := crypto.Keccak256Hash([]byte("Lighthouse"))
	versionHash := crypto.Keccak256Hash([]byte("1"))

	// Define the domain separator
	domainPacked, err := abi.Arguments{
		{Type: bytes32Ty}, // typehash
		{Type: bytes32Ty}, // name hash
		{Type: bytes32Ty}, // version hash
		{Type: uint256Ty}, // chainId
		{Type: addressTy}, // verifying contract
	}.Pack(
		domainTypeHash,
		nameHash,
		versionHash,
		big.NewInt(int64(*l.bidder.Config().LighthouseChainId)),
		common.HexToAddress(*l.bidder.Config().LighthouseContractAddress),
	)
	if err != nil {
		log.Fatalf("Failed to pack domain arguments: %v", err)
	}

	domainSeparator := crypto.Keccak256Hash(domainPacked)

	eip712Prefix := []byte("\x19\x01")
	digest := crypto.Keccak256Hash(
		append(eip712Prefix, append(domainSeparator.Bytes(), structHash.Bytes()...)...),
	)

	signature, err := crypto.Sign(digest.Bytes(), l.bidderPrivateKey)
	if err != nil {
		log.Fatalf("Failed to sign EIP-712 payload: %v", err)
	}

	req := &requests.SubmitBidRequest{
		BidderAddress:   l.bidderAddress,
		AuctionId:       *event.AuctionId,
		BidAmount:       bidAmount,
		MetaTxNonce:     l.bidder.MetaTxNonce(),
		RawTransactions: [][]byte{rawTx},
		TxHashes:        txHashes,
		Signature:       signature,
	}
	if err = l.SendMessage(requests.SubmitBid, req); err != nil {
		return err
	}

	l.bidder.IncreaseNonce()
	l.bidder.IncreaseMetaTxNonce()

	logger.Println("Bid submitted")
	return nil
}

func ConvertToBytes32Array(hashes []common.Hash) [][32]byte {
	result := make([][32]byte, len(hashes))
	for i, h := range hashes {
		result[i] = h
	}
	return result
}

//func (l *LighthouseMessageHandler) createPayload() {
//	payloads := []Lighthouse.ILighthousePayload{
//		{
//			Round:     uint8(3),
//			Bidder:    bidder1,
//			Price:     price1,
//			Nonce:     nonce,
//			TxHashes:  convertedTxHashes,
//			Signature: signature,
//		},
//	}
//	return payloads
//}

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
