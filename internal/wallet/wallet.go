package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"log"
)

type Wallet struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  *ecdsa.PublicKey
}

func GenerateKeyPair() (*ecdsa.PrivateKey, *ecdsa.PublicKey) {
	// Generate a new private key using the P-256 curve (similar to Bitcoin's)
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatalf("Failed to generate key pair: %s", err)
	}

	return privateKey, &privateKey.PublicKey
}

func (w *Wallet) CreateTransaction(receiver *ecdsa.PublicKey, amount float64) *Transaction {
	tx := &Transaction{
		Sender:   w.PublicKey,
		Receiver: receiver,
		Amount:   amount,
	}
	return tx
}
