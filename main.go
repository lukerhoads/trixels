package main

import (
	"log"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/lukerhoads/trixels/server"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	godotenv.Load(".env")
	devChan := make(chan string)

	quit := make(chan struct{})
	twoWeekTicker := time.NewTicker(14 * 24 * time.Hour)

	rpcUrl := os.Getenv("ETH_URL")
	gasAccountPrivateKey := os.Getenv("GAS_ACCOUNT_PK")
	trixelsAuctionHouseAddress := os.Getenv("TRIXELS_AUCTION_ADDRESS")
	dsn := os.Getenv("DB_DSN")

	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		log.Fatal(err)
	}

	trixelsAuctionHouse, err := server.NewTrixelsAuctionHouse(client, trixelsAuctionHouseAddress)
	if err != nil {
		log.Fatal(err)
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.Table("pixels").AutoMigrate(&server.Pixel{})
	// server.EnsureTableExists(db)

	privateKey, err := crypto.HexToECDSA(gasAccountPrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	periodic := server.NewPeriodic(db, privateKey, client, trixelsAuctionHouse, twoWeekTicker, devChan, quit)
	go periodic.Start()

	var app server.App
	app.Initialize(db, devChan)
	// app.InitializePixels()
	app.Run(":8080")
}
