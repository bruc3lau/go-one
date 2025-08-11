package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func main() {
	//ks := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)
	////fmt.Println(ks)
	password := "password"
	//account, err := ks.NewAccount(password)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(account.Address.Hex())

	file := "./wallet/UTC--2025-08-11T17-06-22.517532000Z--049b9bd882bf52785c9bb202f23866ae103a3501"
	readFile, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(readFile))

	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := ks.Import(readFile, password, password)
	if err != nil {
		panic(err)
	}
	fmt.Println(account.Address.Hex())

	if err := os.Remove(file); err != nil {
		log.Fatal(err)
	}
}
