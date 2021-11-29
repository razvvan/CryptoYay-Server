package yay

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/razvvan/CryptoYay-Server/pkg/contracts"
	"github.com/spf13/viper"
)

func Send(to string, org string) error {
	networkURL := viper.GetString("network_url")
	chainID := viper.GetInt("chain_id")
	privateKeyStr := viper.GetString("private_key")
	contractAddress := viper.GetString("contract_address")

	client, err := ethclient.Dial(networkURL)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(int64(chainID)))
	if err != nil {
		log.Fatal(err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	address := common.HexToAddress(contractAddress)
	instance, err := contracts.NewCryptoYay(address, client)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := instance.SafeMint(auth, common.HexToAddress(to), org, "https://foo.json")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex())

	orgs, err := instance.GetOrgs(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Orgs: %#+v", orgs)

	return nil
}
