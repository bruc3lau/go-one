package main

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

func main() {
	//if key, err := crypto.GenerateKey(); err != nil {
	//	fmt.Println(key)
	//} else {
	//	panic(err)
	//}

	key, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(key)

	privateKeyBytes := crypto.FromECDSA(key)
	fmt.Println("privateKeyBytes", privateKeyBytes)

	privateKeyHex := hex.EncodeToString(privateKeyBytes)
	fmt.Println(privateKeyHex)

	privateKeyEncode := hexutil.Encode(privateKeyBytes)
	fmt.Println("privateKeyEncode", privateKeyEncode[2:])

	publicKey := key.Public()
	fmt.Println("publicKey", publicKey)

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	//fmt.Println("publicKeyECDSA", publicKeyECDSA)

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	publicKeyEncode := hexutil.Encode(publicKeyBytes)
	fmt.Println("publicKeyEncode", publicKeyEncode)

	//toString := hex.EncodeToString(publicKey)
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println("address", address)

	//实现获取地址方法，keccak256
	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	encode := hexutil.Encode(hash.Sum(nil)[12:])
	fmt.Println("encode", encode)
}
