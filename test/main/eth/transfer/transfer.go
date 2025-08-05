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

func main2() {
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

// New main function to demonstrate other API calls
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

	fmt.Println("\n--- Demonstrating Other Ethereum API Calls ---")

	// 1. Get Latest Block Number
	getLatestBlockNumber(client)

	// 2. Get Block Details (using the latest block number)
	// For demonstration, let's get details of the latest block.
	// In a real scenario, you might fetch a specific block by its number or hash.
	latestBlock, err := client.BlockByNumber(context.Background(), nil) // nil for latest
	if err != nil {
		log.Printf("Failed to get latest block for details: %v", err)
	} else {
		getBlockDetails(client, latestBlock.Number())
	}

	// 3. Get Account Balance
	// Use a well-known address on Sepolia with some ETH for demonstration
	// This is a random address, replace with one you know has funds on Sepolia if needed.
	demoAddress := common.HexToAddress("0x742d35Cc6634C0532925a3b844Bc454e4438f44e") // Example address
	getAccountBalance(client, demoAddress)

	fmt.Println("\n--- End of API Call Demonstrations ---")
}

// getLatestBlockNumber fetches and prints the latest block number.
func getLatestBlockNumber(client *ethclient.Client) {
	header, err := client.HeaderByNumber(context.Background(), nil) // nil for latest block
	if err != nil {
		log.Printf("❌ Failed to get latest block header: %v", err)
		return
	}
	fmt.Printf("\n--- Latest Block Number ---\n")
	fmt.Printf("    - Latest Block Number: %d\n", header.Number.Uint64())
}

// getBlockDetails fetches and prints details of a specific block.
func getBlockDetails(client *ethclient.Client, blockNumber *big.Int) {
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Printf("❌ Failed to get block details for block %s: %v", blockNumber.String(), err)
		return
	}

	fmt.Printf("\n--- Block Details (Block #%s) ---\n", block.Number().String())
	fmt.Printf("    - Block Hash: %s\n", block.Hash().Hex())
	fmt.Printf("    - Time: %d\n", block.Time())
	fmt.Printf("    - Nonce: %d\n", block.Nonce())
	fmt.Printf("    - Transactions in Block: %d\n", len(block.Transactions()))
	fmt.Printf("    - Gas Used: %d\n", block.GasUsed())
	fmt.Printf("    - Gas Limit: %d\n", block.GasLimit())
	fmt.Printf("    - Miner: %s\n", block.Coinbase().Hex())
}

// getAccountBalance fetches and prints the balance of a given Ethereum address.
func getAccountBalance(client *ethclient.Client, address common.Address) {
	balance, err := client.BalanceAt(context.Background(), address, nil) // nil for latest block
	if err != nil {
		log.Printf("❌ Failed to get balance for address %s: %v", address.Hex(), err)
		return
	}

	// Convert Wei to ETH for display
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(1e18))

	fmt.Printf("\n--- Account Balance ---\n")
	fmt.Printf("    - Address: %s\n", address.Hex())
	fmt.Printf("    - Balance: %s ETH\n", ethValue.String())
}
