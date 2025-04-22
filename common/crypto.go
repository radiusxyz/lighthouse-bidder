package common

import (
	"crypto/ecdsa"
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

func GetSignature(target, hexPrivateKey string) ([]byte, error) {
	privateKey, err := LoadPrivateKey(hexPrivateKey)
	if err != nil {
		return nil, err
	}

	hash := crypto.Keccak256Hash(
		[]byte(target),
	)

	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		return nil, err
	}
	return signature, nil
}
