package server

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"time"

	"gorm.io/gorm"

	"github.com/SkynetLabs/go-skynet/v2"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type PeriodicConfig struct {
	url string
	privateKey string
}

type Periodic struct {
	config *PeriodicConfig
	*gorm.DB
	*TrixelsAuctionHouse
	*ethclient.Client
	twoWeekTicker *time.Ticker
	devChan       chan string
	quit          chan struct{}
}

func NewPeriodic(config *PeriodicConfig, db *gorm.DB, client *ethclient.Client, trixelsAuctionHouse *TrixelsAuctionHouse, twoWeekTicker *time.Ticker, devChan chan string, quit chan struct{}) *Periodic {
	return &Periodic{
		config: config,
		DB:                  db,
		TrixelsAuctionHouse: trixelsAuctionHouse,
		Client:              client,
		twoWeekTicker:       twoWeekTicker,
		devChan:             devChan,
		quit:                quit,
	}
}

func (p *Periodic) Start() {
	for {
		select {
		case <-p.twoWeekTicker.C:
			if err := p.MintAndStartAuction(); err != nil {
				log.Panic(err)
			}
		case v := <-p.devChan:
			if v == "mint" {
				if err := p.MintAndStartAuction(); err != nil {
					log.Panic(err)
				}
			}
		case <-p.quit:
			p.twoWeekTicker.Stop()
		}
	}
}

// CreateImage takes the pixel list and 
func CreateImage(pixels *Pixels) image.Image {
	upLeft := image.Point{0, 0}
	lowRight := image.Point{500, 500}
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})
	for i := 0; i < IMAGE_DIMENSIONS; i++ {
		for j := 0; j < IMAGE_DIMENSIONS; j++ {
			img.Set(i, j, color.Black)
		}
	}

	for i := 0; i < len(*pixels); i++ {
		color, err := ParseHexColor((*pixels)[i].Color)
		if err != nil {
			continue
		}
		img.Set(int((*pixels)[i].X), int((*pixels)[i].Y), color)
	}

	return img
}

// MintAndStartAuction takes the pixel color values, constructs an image, and submits it on-chain.
// It returns an error if something occurs.
func (p *Periodic) MintAndStartAuction() error {
	log.Println("Starting auction...")

	var pixels Pixels
	pixels.GetPixels(p.DB)
	img := CreateImage(&pixels)
	f, err := os.Create(IMAGE_PATH)
	if err != nil {
		return err
	}

	png.Encode(f, img)
	client := skynet.New()
	url, err := client.UploadFile(IMAGE_PATH, skynet.DefaultUploadOptions)
	if err != nil {
		return err
	}

	// Delete image
	err = os.Remove(IMAGE_PATH)
	if err != nil {
		return err
	}

	metadata := NewMetadata(time.Now().String(), url, "Trixels NFT")
	file, err := json.MarshalIndent(metadata, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(METADATA_PATH, file, 0644)
	if err != nil {
		return err
	}

	metaUrl, err := client.UploadFile(METADATA_PATH, skynet.DefaultUploadOptions)
	if err != nil {
		return err
	}

	err = os.Remove(METADATA_PATH)
	if err != nil {
		return err
	}

	var trixel *Trixel
	count := trixel.GetTrixelCount(p.DB)
	newTokenId := uint64(count + 1)
	newTrixel := &Trixel{
		TokenID:     newTokenId,
		MetadataUrl: metaUrl,
	}

	newTrixel.AddTrixel(p.DB)

	// keyedTransactor := p.GenKeyedTransactor()
	// tx, err := p.TrixelsAuctionHouse.StartAuction(keyedTransactor, skyNetId)
	// if err != nil {
	// 	return err
	// }

	// log.Println(tx.Hash().Hex())
	return nil
}

func (p *Periodic) GenKeyedTransactor() *bind.TransactOpts {
	privateKey, err := crypto.HexToECDSA(p.config.privateKey)
	if err != nil {
	  log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := p.Client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := p.Client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice
	return auth
}
