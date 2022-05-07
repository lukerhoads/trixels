package main

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"

	trixels "github.com/lukerhoads/trixels/abigen"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
)

type Trixels struct {
	*ethclient.Client
	*trixels.Trixels
}

func NewTrixels(client *ethclient.Client, addr string) (*Trixels, error) {
	address := common.HexToAddress(addr)
	instance, err := trixels.NewTrixels(address, client)
	if err != nil {
		return nil, err 
	}

	return &Trixels{
		Client: client,
		Trixels: instance,
	}, nil
}

// func (t *Trixels) GetPixels() ([][]Pixel, error) {
// 	colors, err := t.Trixels.Colors(nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Convert to pixel matrix
// 	var pixels [][]Pixel 
// 	for _, c := range colors {
// 		log.Println(c)
// 	}

// 	return pixels, nil
// }

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