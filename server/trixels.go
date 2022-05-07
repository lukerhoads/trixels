package main

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"

	trixels "github.com/lukerhoads/trixels/abigen"
)

type Trixels struct {
	*ethclient.Client
	*trixels.Trixels
}

func NewTrixels(client *ethclient.Client, address string) (*Trixels, error) {
	address := common.HexToAddress(address)
	instance, err := trixels.NewTrixels(address, client)
	if err != nil {
		return nil, err 
	}

	return &Contract{
		Client: client,
		Trixels: instance,
	}, nil
}

func (t *Trixels) GetPixels() ([][]Pixel, error) {
	colors, err := t.Trixels.Colors(nil)
	if err != nil {
		return nil, err
	}

	// Convert to pixel matrix
	var pixels [][]Pixel 
	for _, c := range colors {
		log.Println(c)
	}

	return pixels, nil
}

type TrixelsAuctionHouse struct {
	*ethclient.Client
	*trixels.TrixelsAuctionHouse
}

func NewTrixelsAuctionHouse(client *ethclient.Client, address string) (*TrixelsAuctionHouse, error) {
	address := common.HexToAddress(address)
	instance, err := trixels.NewTrixelsAuctionHouse(address, client)
	if err != nil {
		return nil, err 
	}

	return &TrixelsAuctionHouse{
		Client: client,
		Trixels: instance,
	}, nil
}

func (c *TrixelsAuctionHouse) StartAuction(skyNetId string) error {
	// Attach to the gas account

	// Start auction with given skyNetId
}