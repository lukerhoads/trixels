package main

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"

	trixels "../contracts"
)

type Trixels struct {
	*ethclient.Client
	
}

func NewTrixels(node string, address string) (*Contract, error) {
	client, err := ethclient.Dial(node)
	if err != nil {
		return nil, err 
	}

	return &Contract{
		Client: client,
		Address: common.HexToAddress(address),
	}, nil
}

func (c *Contract) GetPixels() []Pixel {
	
}