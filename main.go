package main

import (
	"fmt"
	"os"
	"log"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/lukerhoads/trixels/server"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	godotenv.Load(".env")

	server.InitLogger()
	if server.Logger == nil {
		log.Panic("Logger not defined")
	}

	dsn := server.PanicOrReturn(os.Getenv("DATABASE_DSN"))
	rpcUrl := server.PanicOrReturn(os.Getenv("ETH_NODE_URL"))
	pk := server.PanicOrReturn(os.Getenv("PROXY_PRIVATE_KEY"))
	auctionHouseAddress := server.PanicOrReturn(os.Getenv("AUCTION_HOUSE_ADDRESS"))

	devChan := make(chan string)
	twoWeekTicker := time.NewTicker(14 * 24 * time.Hour)
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-quit
		server.Logger.Info("\r- Ctrl+C pressed in Terminal, ending...")
		os.Exit(0)
	}()

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		server.Logger.Fatal(fmt.Sprintf("Could not open database connection @ %s", dsn))
	}

	db.Table("pixels").AutoMigrate(&server.Pixel{})
	db.Table("trixels").AutoMigrate(&server.Trixel{})

	auctionHouse, err := server.NewAuctionHouse(rpcUrl, pk, auctionHouseAddress)
	if err != nil {
		server.Logger.Fatal(fmt.Sprintf("Could not open AuctionHouse connection with RPC URL @ %s and address @ %s", rpcUrl, auctionHouseAddress))
	}

	daemon := server.NewDaemon(db, auctionHouse, twoWeekTicker, devChan, quit)
	go daemon.Start()

	var app server.App
	app.Initialize(db, devChan)
	app.Run(":8080")
}
