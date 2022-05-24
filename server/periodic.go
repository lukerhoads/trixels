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
}

type Periodic struct {
	*PeriodicConfig
	*gorm.DB
	*TrixelsAuctionHouse
	*ethclient.Client
	privateKey    *ecdsa.PrivateKey
	twoWeekTicker *time.Ticker
	devChan       chan string
	quit          chan struct{}
}

func NewPeriodic(db *gorm.DB, privateKey *ecdsa.PrivateKey, client *ethclient.Client, trixelsAuctionHouse *TrixelsAuctionHouse, twoWeekTicker *time.Ticker, devChan chan string, quit chan struct{}) *Periodic {
	return &Periodic{
		DB:                  db,
		TrixelsAuctionHouse: trixelsAuctionHouse,
		Client:              client,
		twoWeekTicker:       twoWeekTicker,
		devChan:             devChan,
		quit:                quit,
		privateKey:          privateKey,
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

func (p *Periodic) MintAndStartAuction() error {
	log.Println("Starting auction...")

	var pixels Pixels
	pixels.GetPixels(p.DB)

	// Construct image from pixels
	upLeft := image.Point{0, 0}
	lowRight := image.Point{500, 500}
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})
	for i := 0; i < 500; i++ {
		for j := 0; j < 500; j++ {
			img.Set(i, j, color.Black)
		}
	}

	for i := 0; i < len(pixels); i++ {
		color, err := ParseHexColor(pixels[i].Color)
		if err != nil {
			continue
		}
		img.Set(int(pixels[i].X), int(pixels[i].Y), color)
	}

	path := "image.png"
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	png.Encode(f, img)

	// Put it on SkyNet, get CDN url
	client := skynet.New()
	url, err := client.UploadFile(path, skynet.DefaultUploadOptions)
	if err != nil {
		return err
	}

	// Delete image
	err = os.Remove(path)
	if err != nil {
		return err
	}

	// Mint NFT with metadata URL
	metadata := Metadata{
		Name:        time.Now().String(),
		Image:       url,
		Description: "Trixels NFT",
	}

	metaPath := "metadata.json"
	file, err := json.MarshalIndent(metadata, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(metaPath, file, 0644)
	if err != nil {
		return err
	}

	// Begin auction on NFT
	metaUrl, err := client.UploadFile(metaPath, skynet.DefaultUploadOptions)
	if err != nil {
		return err
	}

	// Delete image
	err = os.Remove(metaPath)
	if err != nil {
		return err
	}

	var trixel *Trixel
	count := trixel.GetTrixelCount(p.DB)

	// Add route53 rule to redirect (add to database)
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
	publicKey := p.privateKey.Public()
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

	auth := bind.NewKeyedTransactor(p.privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice
	return auth
}
