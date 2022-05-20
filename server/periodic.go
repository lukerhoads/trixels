package server

import (
	"log"
	"image"
	"image/color"
	"image/png"
	"time"
	"crypto/ecdsa"
	"context"
	"os"
	"encoding/json"
	"io/ioutil"
	"strings"
	"math/big"
	"fmt"
	"gorm.io/gorm"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/SkynetLabs/go-skynet/v2"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

type Periodic struct {
	*gorm.DB
	*Trixels
	*TrixelsAuctionHouse
	*ethclient.Client
	dayTicker *time.Ticker 
	twoWeekTicker *time.Ticker 
	devChan chan string
	quit chan struct{}
	privateKey *ecdsa.PrivateKey
}

func NewPeriodic(db *gorm.DB, privateKey *ecdsa.PrivateKey, client *ethclient.Client, trixels *Trixels, trixelsAuctionHouse *TrixelsAuctionHouse, dayTicker *time.Ticker, twoWeekTicker *time.Ticker, devChan chan string, quit chan struct{}) *Periodic {
	return &Periodic{
		DB: db,
		Trixels: trixels,
		TrixelsAuctionHouse: trixelsAuctionHouse,
		Client: client,
		dayTicker: dayTicker,
		twoWeekTicker: twoWeekTicker,
		devChan: devChan,
		quit: quit,
		privateKey: privateKey,
	}
}

func (p *Periodic) Start() {
	log.Println("here")
	for {
		select {
		case <-p.dayTicker.C:
			// if err := p.UpdatePixels(); err != nil {
			// 	log.Panic(err)
			// } 
		case <-p.twoWeekTicker.C:
			if err := p.MintAndStartAuction(); err != nil {
				log.Panic(err)
			} 
		// Simply for demo purposes
		case v := <-p.devChan:
			log.Println("pulled off devchan")
			if v == "mint" {
				if err := p.MintAndStartAuction(); err != nil {
					log.Panic(err)
				}
			} else if v == "updatepixels" {
				// if err := p.UpdatePixels(); err != nil {
				// 	log.Panic(err)
				// }
			}
		case <-p.quit:
			p.dayTicker.Stop()
			p.twoWeekTicker.Stop()
		}
	}
}

func (p *Periodic) StartUpdater(addr string) {
	// Setup ethereum event log/listener for an update transaction
	contractAddress := common.HexToAddress(addr)
	query := ethereum.FilterQuery{
		BlockHash: nil,
		FromBlock: nil,
		ToBlock: nil,
		Addresses: []common.Address{contractAddress},
		Topics: [][]common.Hash{{common.HexToHash("0xdfda710484c7dc3d855d2d80e0f7832b0f90247a4a297036e424961e0c590bee")}},
	}

	logs := make(chan types.Log)
	sub, err := p.Client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Panic(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Panic(err)
		case vLog := <-logs:
			// create a commit based off the transaction log
			log.Println(vLog)
		}
	}
}

// func (p *Periodic) UpdatePixels() error {
// 	commits := p.Store.GetDayCommits()

// 	// Write to the contract from a address with ETH
// 	keyedTransactor := p.GenKeyedTransactor()

// 	// Convert commits to new pixels
// 	var pixels []abigen.ITrixelsPixel
// 	for _, c := range commits {
// 		pixels = append(pixels, abigen.ITrixelsPixel{
// 			X: c.X,
// 			Y: c.Y,
// 			Color: [3]byte(c.Color),
// 			LastEditor: common.HexToAddress(c.Address),
// 		})
// 	}

// 	tx, err := p.Trixels.MassChangePixels(keyedTransactor, pixels)
// 	if err != nil {
// 		return err 
// 	}

// 	log.Println(tx.Hash().Hex())
// 	return nil
// }

func (p *Periodic) MintAndStartAuction() error {
	var pixels []Pixel
	p.DB.Find(&pixels)

	// Construct image from pixels
	upLeft := image.Point{0, 0}
	lowRight := image.Point{30, 30}
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})
	for i := 0; i < 30; i++ {
		for j := 0; j < 30; j++ {
			img.Set(i, j, color.Black)
		}
	}

	for i := 0; i < len(pixels); i++ {
		color, err := ParseHexColor(pixels[i].Color)
		if err != nil {
			return err
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
		Name: time.Now().String(),
		Image: url,
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

	skyNetId := metaUrl[strings.LastIndex(metaUrl, "/")+1:]
	keyedTransactor := p.GenKeyedTransactor()
	tx, err := p.TrixelsAuctionHouse.StartAuction(keyedTransactor, skyNetId)
	if err != nil {
		return err 
	}

	log.Println(tx.Hash().Hex())
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

func ParseHexColor(s string) (c color.RGBA, err error) {
    c.A = 0xff
    switch len(s) {
    case 7:
        _, err = fmt.Sscanf(s, "#%02x%02x%02x", &c.R, &c.G, &c.B)
    case 4:
        _, err = fmt.Sscanf(s, "#%1x%1x%1x", &c.R, &c.G, &c.B)
        // Double the hex digits:
        c.R *= 17
        c.G *= 17
        c.B *= 17
    default:
        err = fmt.Errorf("invalid length, must be 7 or 4")
    }

    return
}