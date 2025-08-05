package ethapi

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// GetTransactionByHash fetches and prints details of a transaction by its hash.
func GetTransactionByHash(client *ethclient.Client, txHash common.Hash) {
	fmt.Printf("\n--- Transaction Details (Hash: %s) ---\n", txHash.Hex())
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Printf("❌ Failed to get transaction by hash %s: %v", txHash.Hex(), err)
		return
	}

	fmt.Printf("    - Pending: %t\n", isPending)
	if tx != nil {
		fmt.Printf("    - Value: %s Wei\n", tx.Value().String())
		fmt.Printf("    - Gas Price: %s Wei\n", tx.GasPrice().String())
		fmt.Printf("    - Gas: %d\n", tx.Gas())
		fmt.Printf("    - Nonce: %d\n", tx.Nonce())
		fmt.Printf("    - To: %s\n", tx.To().Hex())
		// You can get more details like 'From' address by signing the transaction
		// or by getting the transaction receipt.
	} else {
		fmt.Printf("    - Transaction not found.\n")
	}
}

// GetTransactionReceipt fetches and prints the receipt of a transaction by its hash.
func GetTransactionReceipt(client *ethclient.Client, txHash common.Hash) {
	fmt.Printf("\n--- Transaction Receipt (Hash: %s) ---\n", txHash.Hex())
	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		log.Printf("❌ Failed to get transaction receipt for hash %s: %v", txHash.Hex(), err)
		return
	}

	if receipt != nil {
		fmt.Printf("    - Status: %d (1=Success, 0=Fail)\n", receipt.Status)
		fmt.Printf("    - Block Number: %s\n", receipt.BlockNumber.String())
		fmt.Printf("    - Gas Used: %d\n", receipt.GasUsed)
		fmt.Printf("    - Cumulative Gas Used: %d\n", receipt.CumulativeGasUsed)
		fmt.Printf("    - Contract Address: %s\n", receipt.ContractAddress.Hex())
		fmt.Printf("    - Logs Count: %d\n", len(receipt.Logs))
	} else {
		fmt.Printf("    - Transaction receipt not found (might be pending or invalid hash).\n")
	}
}
