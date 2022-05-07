package main

import (
	"time"
	"os"
	"log"
)

func main() {
	dayTicker := time.NewTicker(24 * time.Hour)
	twoWeekTicker := time.NewTicker(14 * 24 * time.Hour)
	devChan := make(chan string)

	rpcUrl := os.Getenv("ETH_URL")
	gasAccountPrivateKey := os.Getenv("GAS_ACCOUNT_PK")
	trixelsAddress := os.Getenv("TRIXELS_ADDRESS")
	trixelsAuctionHouseAddress := os.Getenv("TRIXELS_AUCTION_ADDRESS")
	dsn := os.Getenv("DB_DSN")

	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		log.Fatal(err)
	}

	trixels := NewTrixels(client, trixelsAddress)
	trixelsAuctionHouse := NewTrixelsAuctionHouse(client, trixelsAuctionHouseAddress)

	store, err := NewStore(dsn)
	if err != nil {
		log.Panic(err)
	}

	privateKey, err := crypto.HexToECDSA(gasAccountPrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	periodic := NewPeriodic(store, privateKey, client, trixels, trixelsAuctionHouse, dayTicker, twoWeekTicker, devChan)
	go periodic.Start()
	go periodic.StartUpdater()

	router := NewRouter(store, devChan)
	router.Start()
}