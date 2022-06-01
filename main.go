package main

import (
	"log"
	"os"
	"time"

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

	dsn := os.Getenv("DB_DSN")
	rpcUrl := os.Getenv("ETH_URL")
	pk := os.Getenv("GAS_ACCOUNT_PK")
	auctionHouseAddress := os.Getenv("AUCTION_HOUSE_ADDRESS")

	auctionHouse, err := server.NewAuctionHouse(rpcUrl, pk, auctionHouseAddress)
	if err != nil {
		log.Fatal(err)
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.Table("pixels").AutoMigrate(&server.Pixel{})
	periodic := server.NewDaemon(db, auctionHouse, twoWeekTicker, devChan, quit)
	go periodic.Start()

	var app server.App
	app.Initialize(db, devChan)
	app.Run(":8080")
}
