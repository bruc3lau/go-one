package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main1() {
	// --- 1. Configuration (Replace with your details) ---

	// You can get a free ID from https://infura.io
	infuraProjectID := "YOUR_INFURA_PROJECT_ID"

	// WARNING: Never hardcode private keys in production code.
	// Use environment variables or a secure vault for real applications.
	// This private key must correspond to an account that has Sepolia test ETH.
	senderPrivateKeyHex := "YOUR_SENDER_PRIVATE_KEY"

	// --- 2. Connect to the Sepolia Test Network ---
	infuraURL := fmt.Sprintf("https://sepolia.infura.io/v3/%s", infuraProjectID)
	client, err := ethclient.Dial(infuraURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	fmt.Println("✅ Successfully connected to Sepolia network!")

	// --- 3. Call the transfer function ---
	err = sendTestTransaction(client, senderPrivateKeyHex)
	if err != nil {
		log.Fatalf("❌ Transaction failed: %v", err)
	}
}

func main() {
	// --- 1. Configuration (Replace with your details) ---

	// You can get a free ID from https://infura.io
	infuraProjectID := os.Getenv("INFURA_PROJECT_ID")

	// WARNING: Never hardcode private keys in production code.
	// Use environment variables or a secure vault for real applications.
	// This private key must correspond to an account that has Sepolia test ETH.
	senderPrivateKeyHex := os.Getenv("SENDER_PRIVATE_KEY")
	senderPrivateKeyHex = strings.TrimPrefix(senderPrivateKeyHex, "0x")

	// --- 2. Connect to the Sepolia Test Network ---
	infuraURL := fmt.Sprintf("https://sepolia.infura.io/v3/%s", infuraProjectID)
	client, err := ethclient.Dial(infuraURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	fmt.Println("✅ Successfully connected to Sepolia network!")

	// --- 3. Call the transfer function ---
	err = sendTestTransaction(client, senderPrivateKeyHex)
	if err != nil {
		log.Fatalf("❌ Transaction failed: %v", err)
	}
}

// sendTestTransaction prepares and sends a small amount of test ETH.
func sendTestTransaction(client *ethclient.Client, senderPrivateKeyHex string) error {
	// --- Load Sender's Account from Private Key ---
	privateKey, err := crypto.HexToECDSA(senderPrivateKeyHex)
	if err != nil {
		return fmt.Errorf("error parsing private key: %w", err)
	}

	// Derive the public address from the private key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Printf("    - Sender Address: %s\n", fromAddress.Hex())

	// --- Define Recipient and Amount ---
	toAddress := common.HexToAddress("0x5F6583Cc3e7ae857910915841E093BB498dC0fD6")
	// Let's send a very small amount: 0.001 ETH
	// 1 ETH = 10^18 Wei. So, 0.001 ETH = 10^15 Wei.
	amount := big.NewInt(1000000000000000) // 0.001 ETH in Wei

	// --- Get Required Transaction Details from the Network ---
	// Get the next nonce for the sender's account (to prevent replay attacks)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return fmt.Errorf("failed to get nonce: %w", err)
	}

	// Get the suggested gas price from the network
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get gas price: %w", err)
	}

	// Set the gas limit for a standard ETH transfer
	gasLimit := uint64(21000)

	// --- Create and Sign the Transaction ---
	// Create the transaction object with all the details
	tx := types.NewTransaction(nonce, toAddress, amount, gasLimit, gasPrice, nil)

	// Get the network chain ID for signing (Sepolia's ID is 11155111)
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get chain ID: %w", err)
	}

	// Sign the transaction with the sender's private key to prove ownership
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return fmt.Errorf("failed to sign transaction: %w", err)
	}

	// --- Broadcast the Transaction to the Network ---
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return fmt.Errorf("failed to send transaction: %w", err)
	}

	// --- Print the Result ---
	fmt.Printf("✅ Transaction successfully sent!\n")
	fmt.Printf("    - Recipient: %s\n", toAddress.Hex())
	// Convert Wei back to ETH for display
	ethValue := new(big.Float).Quo(new(big.Float).SetInt(amount), big.NewFloat(1e18))
	fmt.Printf("    - Amount: %s ETH\n", ethValue.String())
	fmt.Printf("    - Tx Hash: %s\n", signedTx.Hash().Hex())
	fmt.Printf("    - View on Etherscan: https://sepolia.etherscan.io/tx/%s\n", signedTx.Hash().Hex())

	return nil
}
