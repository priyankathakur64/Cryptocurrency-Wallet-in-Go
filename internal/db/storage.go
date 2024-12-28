package db

import (
	"encoding/json"
	"os"
)

func SaveWalletData(wallet wallet) error {
	file, err := os.OpenFile("wallet_data.json", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(wallet)
}

func LoadWalletData() (*Wallet, error) {
	file, err := os.Open("wallet_data.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var wallet Wallet
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&wallet)
	return &wallet, err
}
