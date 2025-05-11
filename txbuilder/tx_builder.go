package txbuilder

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

type TxBuilder struct {
	//rpcProvider RpcProvider
	ethClient *ethclient.Client
}

func New(url string) (*TxBuilder, error) {
	client, err := ethclient.DialContext(context.Background(), url)
	if err != nil {
		return nil, err
	}

	return &TxBuilder{
		ethClient: client,
	}, nil
}

func (t *TxBuilder) GetSignedTransaction(privateKeyHex string, toAddress common.Address) (string, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatalf("invalid private key: %v", err)
	}

	// 3. 주소, nonce 조회
	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	fromAddress := crypto.PubkeyToAddress(*publicKey)

	nonce, err := t.ethClient.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalf("failed to get nonce: %v", err)
	}

	// 4. 트랜잭션 파라미터 설정
	value := big.NewInt(10000000000000000)                             // 0.01 ETH
	gasLimit := uint64(21000)                                          // 기본 전송
	gasPrice, err := t.ethClient.SuggestGasPrice(context.Background()) // legacy tx
	if err != nil {
		log.Fatalf("failed to suggest gas price: %v", err)
	}

	// 5. 트랜잭션 생성 (legacy 트랜잭션 예시)
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)

	// 6. 체인 ID 확인 (ex: 1=mainnet, 5=Goerli, 1337=anvil)
	chainID, err := t.ethClient.NetworkID(context.Background())
	if err != nil {
		log.Fatalf("failed to get chain ID: %v", err)
	}

	// 7. 서명 및 RLP 인코딩
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatalf("failed to sign tx: %v", err)
	}

	// 8. rawTransaction 생성
	rawTxBytes, err := signedTx.MarshalBinary()
	if err != nil {
		log.Fatalf("failed to encode tx: %v", err)
	}

	rawTxHex := "0x" + hex.EncodeToString(rawTxBytes)
	fmt.Println("Raw Transaction:", rawTxHex)
	return rawTxHex, nil
}
