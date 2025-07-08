package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

func main1() {
	// WARNING: Do NOT hardcode your real mnemonic phrase in source code.
	// This is for demonstration purposes only. Use environment variables or a secure vault.
	mnemonic := "tag volcano eight thank tide danger coast health above argue embrace heavy" // Replace with your mnemonic phrase

	// Create a new wallet from the mnemonic phrase
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		log.Fatalf("Failed to create wallet from mnemonic: %v", err)
	}

	// Define the standard Ethereum derivation path for the first account (index 0)
	// m / purpose' / coin_type' / account' / change / address_index
	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")

	// Derive the account at the specified path
	account, err := wallet.Derive(path, true) // Set `true` to get the address
	if err != nil {
		log.Fatalf("Failed to derive account: %v", err)
	}

	// Get the private key
	privateKey, err := wallet.PrivateKey(account)
	if err != nil {
		log.Fatalf("Failed to get private key: %v", err)
	}

	// Get the public address
	address := account.Address.Hex()

	// Convert the private key to a hex string for display
	privateKeyHex := hexutil.Encode(crypto.FromECDSA(privateKey))

	fmt.Println("✅ Successfully derived keys from mnemonic!")
	fmt.Printf("   - Address:       %s\n", address)
	fmt.Printf("   - Private Key:   %s\n", privateKeyHex)
}

func main() {
	mnemonic := os.Getenv("mnemonic")

	// WARNING: Do NOT hardcode your real mnemonic phrase in source code.
	// This is for demonstration purposes only. Use environment variables or a secure vault.
	//mnemonic := "tag volcano eight thank tide danger coast health above argue embrace heavy" // Replace with your mnemonic phrase

	// Create a new wallet from the mnemonic phrase
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		log.Fatalf("Failed to create wallet from mnemonic: %v", err)
	}

	// Define the standard Ethereum derivation path for the first account (index 0)
	// m / purpose' / coin_type' / account' / change / address_index
	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")

	// Derive the account at the specified path
	account, err := wallet.Derive(path, true) // Set `true` to get the address
	if err != nil {
		log.Fatalf("Failed to derive account: %v", err)
	}

	// Get the private key
	privateKey, err := wallet.PrivateKey(account)
	if err != nil {
		log.Fatalf("Failed to get private key: %v", err)
	}

	// Get the public address
	address := account.Address.Hex()

	// Convert the private key to a hex string for display
	privateKeyHex := hexutil.Encode(crypto.FromECDSA(privateKey))

	fmt.Println("✅ Successfully derived keys from mnemonic!")
	fmt.Printf("   - Address:       %s\n", address)
	fmt.Printf("   - Private Key:   %s\n", privateKeyHex)
}
