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
	dsn := os.Getenv("DB_DSN")
	store, err := NewStore(dsn)
	if err != nil {
		log.Panic(err)
	}

	updater := NewUpdater(store)
	go updater.Start()

	periodic := NewPeriodic(store, dayTicker, twoWeekTicker, devChan)
	go periodic.Start()

	router := NewRouter(store, devChan)
	router.Start()
}