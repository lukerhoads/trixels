package server

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	aux "github.com/lukerhoads/trixels/abigen/AuctionHouse"
)

type AuctionHouse struct {
	privateKey string
	*ethclient.Client
	*aux.AuctionHouse
}

func NewAuctionHouse(rpcUrl string, privateKey string) (*AuctionHouse, error) {
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		return nil, err
	}

	return &AuctionHouse{
		privateKey: privateKey,
		Client:     client,
	}, nil
}

func (a *AuctionHouse) StartAuction() (uint64, error) {
	_, err := a.AuctionHouse.StartAuction(genKeyedTransactor(a.Client, a.privateKey))
	return 0, err
}

func genKeyedTransactor(client *ethclient.Client, pKey string) *bind.TransactOpts {
	privateKey, err := crypto.HexToECDSA(pKey)
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
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice
	return auth
}
