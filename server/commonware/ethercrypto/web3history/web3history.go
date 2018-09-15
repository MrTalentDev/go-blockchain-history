package web3history

import (
	"fmt"
	"log"
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"./contracts" // for demo
	"github.com/ethereum/go-ethereum/crypto"
	"crypto/ecdsa"
	"math/big"
	"os"
)

var (
	PVT_KEY          = os.Getenv("PVT_KEY")
	CONTRACT_ADDRESS = os.Getenv("CONTRACT_ADDRESS")
)

func SendNewRootHash(rootHash []byte) {
	client, err := ethclient.Dial("https://rinkeby.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(PVT_KEY)
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

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	address := common.HexToAddress(CONTRACT_ADDRESS)
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}
	var hash [32]byte
	copy(hash[:], rootHash)
	fmt.Println(hash)
	_, err = instance.SetRootHash(auth, hash)
	if err != nil {
		log.Fatal(err)
	}
}
