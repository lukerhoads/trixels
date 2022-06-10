package server

import (
	"fmt"
	"encoding/json"
	"image/png"
	"io/ioutil"
	"os"
	"time"
	"strings"

	"gorm.io/gorm"
	"github.com/SkynetLabs/go-skynet/v2"
)

// Daemon is the process running in the background, waiting on dev/time channels.
type Daemon struct {
	*AuctionHouse
	*gorm.DB
	twoWeekTicker *time.Ticker
	devChan       chan string
	quit          chan os.Signal
}

func NewDaemon(db *gorm.DB, auctionHouse *AuctionHouse, twoWeekTicker *time.Ticker, devChan chan string, quit chan os.Signal) *Daemon {
	return &Daemon{
		DB:            db,
		AuctionHouse:  auctionHouse,
		twoWeekTicker: twoWeekTicker,
		devChan:       devChan,
		quit:          quit,
	}
}

// Start will start the daemon, listening on the channels it was given during initialization.
func (p *Daemon) Start() {
	for {
		select {
		case <-p.twoWeekTicker.C:
			Logger.Info("🔨 Starting auction from twoWeekTicker...")
			if err := p.MintAndStartAuction(); err != nil {
				Logger.Error(err.Error())
			}
		case v := <-p.devChan:
			Logger.Info("🔨 Starting auction from devChan...")
			if v == "mint" {
				if err := p.MintAndStartAuction(); err != nil {
					Logger.Error(err.Error())
				}
			}
		case <-p.quit:
			Logger.Info("Stopping daemon...")
			p.twoWeekTicker.Stop()
			return
		}
	}
}

// MintAndStartAuction takes the pixel color values, constructs an image, and submits it on-chain.
// It returns an error.
func (p *Daemon) MintAndStartAuction() error {
	var pixels Pixels
	pixels.GetPixels(p.DB)
	img := CreateImage(&pixels)
	f, err := os.Create(IMAGE_PATH)
	if err != nil {
		return fmt.Errorf("Error creating image file: %s", err)
	}

	png.Encode(f, img)
	client := skynet.New()
	url, err := client.UploadFile(IMAGE_PATH, skynet.DefaultUploadOptions)
	if err != nil {
		return fmt.Errorf("Error uploading image file to Skynet: %s", err)
	}

	err = os.Remove(IMAGE_PATH)
	if err != nil {
		return fmt.Errorf("Error removing image file: %s", err)
	}

	lastPart := strings.Split(url, "//")[1]
	newUrl := strings.Join([]string{"https://siasky.net/", lastPart}, "")
	metadata := NewMetadata(time.Now().String(), newUrl, "Trixels NFT")
	file, _ := json.MarshalIndent(metadata, "", " ")
	err = ioutil.WriteFile(METADATA_PATH, file, 0644)
	if err != nil {
		return fmt.Errorf("Error writing to metadata file: %s", err)
	}

	metaUrl, err := client.UploadFile(METADATA_PATH, skynet.DefaultUploadOptions)
	if err != nil {
		return fmt.Errorf("Error uploading metadata url to Skynet: %s", err)
	}

	err = os.Remove(METADATA_PATH)
	if err != nil {
		return fmt.Errorf("Error removing metadata file: %s", err)
	}

	tokenID, err := p.AuctionHouse.StartAuction()
	if err != nil {
		return fmt.Errorf("Error starting auction: %s", err)
	}

	metaLastPart := strings.Split(metaUrl, "//")[1]
	newMetaUrl := strings.Join([]string{"https://siasky.net/", metaLastPart}, "")
	newTrixel := &Trixel{
		TokenID:     uint(tokenID),
		MetadataUrl: newMetaUrl,
		live: true,
	}

	newTrixel.AddLiveTrixel(p.DB)
	Logger.Info("🔨 Started auction...")
	return nil
}

func (p *Daemon) EndAuction() error {
	var newTrixel Trixel
	found := newTrixel.GetLiveTrixel(p.DB)
	if !found {
		return nil
	}

	newTrixel.DeleteLiveTrixel(p.DB)
	newTrixel.live = false
	newTrixel.AddTrixel(p.DB)
	return nil
}
