package server

import (
	"encoding/json"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"time"
	"strings"

	"gorm.io/gorm"

	"github.com/SkynetLabs/go-skynet/v2"
)

type Daemon struct {
	*AuctionHouse
	*gorm.DB
	twoWeekTicker *time.Ticker
	devChan       chan string
	quit          chan struct{}
}

func NewDaemon(db *gorm.DB, auctionHouse *AuctionHouse, twoWeekTicker *time.Ticker, devChan chan string, quit chan struct{}) *Daemon {
	return &Daemon{
		DB:            db,
		AuctionHouse:  auctionHouse,
		twoWeekTicker: twoWeekTicker,
		devChan:       devChan,
		quit:          quit,
	}
}

func (p *Daemon) Start() {
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

// CreateImage takes pixel hex values and converts them to an image.
// It returns the image.
func CreateImage(pixels *Pixels) image.Image {
	upLeft := image.Point{0, 0}
	lowRight := image.Point{IMAGE_DIMENSIONS, IMAGE_DIMENSIONS}
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
func (p *Daemon) MintAndStartAuction() error {
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

	lastPart := strings.Split(url, "//")[1]
	newUrl := strings.Join([]string{"https://siasky.net/", lastPart}, "")
	metadata := NewMetadata(time.Now().String(), newUrl, "Trixels NFT")
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

	tokenID, err := p.AuctionHouse.StartAuction()
	if err != nil {
		return err
	}

	metaLastPart := strings.Split(metaUrl, "//")[1]
	newMetaUrl := strings.Join([]string{"https://siasky.net/", metaLastPart}, "")
	newTrixel := &Trixel{
		TokenID:     uint(tokenID),
		MetadataUrl: newMetaUrl,
	}

	newTrixel.AddTrixel(p.DB)
	return nil
}
