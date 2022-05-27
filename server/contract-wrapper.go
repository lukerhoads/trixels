package server

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type AuctionHouse struct {
	privateKey string
	*ethclient.Client
	*TrixelsAuctionHouse
}

func NewAuctionHouse(rpcUrl string, privateKey string) (*AuctionHouse, error) {
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		return nil, err
	}

	return &AuctionHouse{
		privateKey: privateKey,
		Client: client,
	}
}

func (a *AuctionHouse) StartAuction(int skyNetID) error {
	_, err := a.TrixelsAuctionHouse.StartAuction(genKeyedTransactor(a.Client, a.privateKey), skyNetID)
	return err
}

func genKeyedTransactor(client *ethclient.Client, privateKey string) *bind.TransactOpts {
	privateKey, err := crypto.HexToECDSA(p.config.privateKey)
	if err != nil {
	  log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := p.Client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := p.Client.SuggestGasPrice(context.Background())
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