package common

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"strings"
)

func LoadPrivateKey(hexPrivateKey string) (*ecdsa.PrivateKey, error) {
	cleanKey := strings.TrimPrefix(hexPrivateKey, "0x")

	privateKey, err := crypto.HexToECDSA(cleanKey)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func GetSignature(target string, timestamp uint64, hexPrivateKey string) ([]byte, error) {
	privateKey, err := LoadPrivateKey(hexPrivateKey)
	if err != nil {
		return nil, err
	}

	message := fmt.Sprintf("%s|%d", target, timestamp)

	hash := crypto.Keccak256Hash(
		[]byte(message),
	)

	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		return nil, err
	}
	return signature, nil
}
