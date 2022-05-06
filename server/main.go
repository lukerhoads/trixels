package main

import (
	"time"
	"os"
	"log"
)

func main() {
	dayTicker := time.NewTicker(24 * time.Hour)
	twoWeekTicker := time.NewTicker(14 * 24 * time.Hour)

	rpcUrl := os.Getenv("ETH_ENDPOINT")
	gasAccountPrivateKey := os.Getenv("GAS_ACCOUNT_PK")
	dsn := os.Getenv("DSN")
	store, err := NewStore(dsn)
	if err != nil {
		log.Panic(err)
	}

	listener := NewListener(store)
	go listener.Start()

	router := NewRouter(store)
	router.Start()
}