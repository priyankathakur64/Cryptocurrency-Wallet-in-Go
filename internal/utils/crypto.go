package utils

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"log"
)

func GenerateKeyPair() (*ecdsa.PrivateKey, *ecdsa.PublicKey) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatalf("Failed to generate key pair: %s", err)
	}

	return privateKey, &privateKey.PublicKey
}
