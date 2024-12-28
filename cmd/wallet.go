package main

import (
	"cryptocurrency-wallet/internal/wallet"
	"fmt"
	"log"
)

func main() {
	// Generate a new key pair for the sender (this is the wallet's key)
	privateKey, publicKey := wallet.GenerateKeyPair()

	// Create a new wallet with the generated keys
	myWallet := &wallet.Wallet{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
	}

	fmt.Println("Your public key:", publicKey)
	fmt.Println("Your private key:", privateKey)

	// Create a new key pair for the receiver (this is for transaction testing)
	_, receiverPublicKey := wallet.GenerateKeyPair()

	// Create a transaction from sender to receiver (sending some amount)
	tx := myWallet.CreateTransaction(receiverPublicKey, 50.0)

	// Sign the transaction with the sender's private key
	signature, err := wallet.SignTransaction(myWallet.PrivateKey, tx)
	if err != nil {
		log.Fatalf("Error signing transaction: %s", err)
	}

	fmt.Println("Transaction Signed with Signature:", signature)

	// Verify the transaction's validity
	isValid := wallet.VerifyTransaction(tx, signature)
	fmt.Println("Is the transaction valid?", isValid)
}
