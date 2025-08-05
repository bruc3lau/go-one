package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"go-one/internal/ethapi"
)

func main() {
	infuraProjectID := os.Getenv("INFURA_PROJECT_ID")
	if infuraProjectID == "" {
		log.Fatal("INFURA_PROJECT_ID environment variable not set.")
	}

	infuraURL := fmt.Sprintf("https://sepolia.infura.io/v3/%s", infuraProjectID)
	client, err := ethclient.Dial(infuraURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	fmt.Println("✅ Successfully connected to Sepolia network!")

	fmt.Println("\n--- Demonstrating New Ethereum API Calls from ethapi package ---")

	// IMPORTANT: Replace with a real transaction hash from Sepolia network
	// You can find one on Sepolia Etherscan: https://sepolia.etherscan.io/
	// For example, a recent transaction hash.
	// This is just a placeholder, it might not exist or be valid.
	// A transaction hash from a successful transfer is ideal.
	txHashStr := "0x5f6583cc3e7ae857910915841e093bb498dc0fd65f6583cc3e7ae857910915841e093bb498dc0fd6" // Placeholder
	// A more realistic example (replace with a real one from Sepolia Etherscan):
	// txHashStr := "0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"

	txHash := common.HexToHash(txHashStr)

	// 1. Get Transaction Details by Hash
	ethapi.GetTransactionByHash(client, txHash)

	// 2. Get Transaction Receipt by Hash
	ethapi.GetTransactionReceipt(client, txHash)

	fmt.Println("\n--- End of New API Call Demonstrations ---")
}
