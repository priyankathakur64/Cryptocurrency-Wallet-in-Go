package wallet

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
)

type Transaction struct {
	Sender    *ecdsa.PublicKey
	Receiver  *ecdsa.PublicKey
	Amount    float64
	Signature []byte
}

func SignTransaction(privateKey *ecdsa.PrivateKey, tx *Transaction) ([]byte, error) {
	txData := fmt.Sprintf("%s%s%f", tx.Sender, tx.Receiver, tx.Amount)
	hash := sha256.Sum256([]byte(txData))

	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
	if err != nil {
		return nil, err
	}

	signature := append(r.Bytes(), s.Bytes()...)
	return signature, nil
}

func VerifyTransaction(tx *Transaction, signature []byte) bool {
	r := new(big.Int)
	s := new(big.Int)
	r.SetBytes(signature[:len(signature)/2])
	s.SetBytes(signature[len(signature)/2:])

	hash := sha256.Sum256([]byte(fmt.Sprintf("%s%s%f", tx.Sender, tx.Receiver, tx.Amount)))
	valid := ecdsa.Verify(tx.Sender, hash[:], r, s)
	return valid
}
