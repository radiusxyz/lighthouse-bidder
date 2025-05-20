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
	"strings"
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

func (t *TxBuilder) GetSignedTransaction(privateKeyHex string, toAddress common.Address, myCurrentAuctionTxCount uint64) (string, error) {
	privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(privateKeyHex, "0x"))
	if err != nil {
		log.Fatalf("invalid private key: %v", err)
	}

	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	fromAddress := crypto.PubkeyToAddress(*publicKey)

	nonce, err := t.ethClient.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalf("failed to get nonce: %v", err)
	}
	fmt.Println("Nonce:", nonce)
	fmt.Println("myCurrentAuctionTxCount:", myCurrentAuctionTxCount)
	value := big.NewInt(10000000000000000)
	gasLimit := uint64(21000)
	gasPrice, err := t.ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("failed to suggest gas price: %v", err)
	}
	fmt.Println("Gas Price:", gasPrice)
	tx := types.NewTransaction(nonce+myCurrentAuctionTxCount, toAddress, value, gasLimit, gasPrice, nil)

	chainID, err := t.ethClient.ChainID(context.Background())
	if err != nil {
		log.Fatalf("failed to get chain ID: %v", err)
	}
	fmt.Println("Chain ID:", chainID)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatalf("failed to sign tx: %v", err)
	}

	rawTxBytes, err := signedTx.MarshalBinary()
	if err != nil {
		log.Fatalf("failed to encode tx: %v", err)
	}

	rawTxHex := "0x" + hex.EncodeToString(rawTxBytes)
	fmt.Println("Raw Transaction:", rawTxHex)
	return rawTxHex, nil
}
