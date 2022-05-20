package server

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"

	trixels "github.com/lukerhoads/trixels/abigen"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
)


type TrixelsAuctionHouse struct {
	*ethclient.Client
	*trixels.TrixelsAuctionHouse
}

func NewTrixelsAuctionHouse(client *ethclient.Client, addr string) (*TrixelsAuctionHouse, error) {
	address := common.HexToAddress(addr)
	instance, err := trixels.NewTrixelsAuctionHouse(address, client)
	if err != nil {
		return nil, err 
	}

	return &TrixelsAuctionHouse{
		Client: client,
		TrixelsAuctionHouse: instance,
	}, nil
}

func (c *TrixelsAuctionHouse) StartAuction(keyedTransactor *bind.TransactOpts, skyNetId string) (*types.Transaction, error) {
	// Attach to the gas account
	var data [23]byte 
	copy(data[:], []byte(skyNetId))
	return c.TrixelsAuctionHouse.StartAuction(keyedTransactor, data)
	// Start auction with given skyNetId
}