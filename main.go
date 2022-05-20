package main

import (
	// "time"
	"os"
	"log"

	// "github.com/ethereum/go-ethereum/ethclient"
	// "github.com/ethereum/go-ethereum/crypto"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"github.com/lukerhoads/trixels/server"
)

func main() {
	godotenv.Load(".env")
	devChan := make(chan string)

	// quit := make(chan struct{})
	// dayTicker := time.NewTicker(24 * time.Hour)
	// twoWeekTicker := time.NewTicker(14 * 24 * time.Hour)
	
	// gasAccountPrivateKey := os.Getenv("GAS_ACCOUNT_PK")
	// trixelsAddress := os.Getenv("TRIXELS_ADDRESS")
	// trixelsAuctionHouseAddress := os.Getenv("TRIXELS_AUCTION_ADDRESS")

	// rpcUrl := os.Getenv("ETH_URL")
	dsn := os.Getenv("DB_DSN")

	// client, err := ethclient.Dial(rpcUrl)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// trixels, err := server.NewTrixels(client, trixelsAddress)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// trixelsAuctionHouse, err := server.NewTrixelsAuctionHouse(client, trixelsAuctionHouseAddress)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.Table("pixels").AutoMigrate(&server.Pixel{})
	// server.EnsureTableExists(db)

	// pixel := Pixel{
	// 	ID: 0,
	// 	X: 0,
	// 	Y: 0,
	// 	LastAddress: "",
	// 	Color: "#000",
	// }
	// result := db.Create(&pixel)
	// if result.Error != nil {
	// 	log.Fatal(result.Error)
	// }

	// privateKey, err := crypto.HexToECDSA(gasAccountPrivateKey)
	// if err != nil {
	// 	log.Fatal(err)
	// }


	// periodic := server.NewPeriodic(db, privateKey, client, trixels, trixelsAuctionHouse, dayTicker, twoWeekTicker, devChan, quit)
	// go periodic.Start()

	var app server.App
	app.Initialize(db, devChan)
	// app.InitializePixels()
	app.Run(":8080")

	
	// go periodic.StartUpdater(trixelsAddress)
}