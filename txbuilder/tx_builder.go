package txbuilder

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

type TxBuilder struct {
	rpcNodeHttpClient *ethclient.Client
}

func New(rpcNodeHttpClient *ethclient.Client, url string) (*TxBuilder, error) {
	rpcNodeHttpClient, err := ethclient.DialContext(context.Background(), url)
	if err != nil {
		return nil, err
	}

	return &TxBuilder{
		rpcNodeHttpClient: rpcNodeHttpClient,
	}, nil
}

func (t *TxBuilder) GetSignedTransaction(privateKey *ecdsa.PrivateKey, toAddress common.Address, nonce uint64) (*types.Transaction, error) {
	value := big.NewInt(10000000000000000)
	gasLimit := uint64(21000)
	gasPrice, err := t.rpcNodeHttpClient.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("failed to suggest gas price: %v", err)
	}

	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)

	chainID, err := t.rpcNodeHttpClient.ChainID(context.Background())
	if err != nil {
		log.Fatalf("failed to get chain ID: %v", err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatalf("failed to sign tx: %v", err)
	}

	return signedTx, nil
}
