package main

import (
	"image"
	"image/color"
	"image/png"

	"github.com/SkynetLabs/go-skynet/v2"
)

type Periodic struct {
	*Store
	*Trixels
	dayTicker *time.Ticker 
	twoWeekTicker *time.Ticker 
	quit chan struct{}
}

func NewPeriodic(store *Store, dayTicker *time.Ticker, twoWeekTicker *time.Ticker, quit chan struct{}) *Periodic {
	return &Periodic{
		Store: store,
		dayTicker,
		twoWeekTicker,
		quit,
	}
}

func (p *Periodic) Start() {
	for {
		select {
		case <-p.dayTicker:
			p.UpdatePixels()
		case <-p.twoWeekTicker:
			err := p.MintAndStartAuction()
			if err != nil {
				log.Panic(err)
			}
		case <-quit:
			p.dayTicker.Stop()
			p.twoWeekTicker.Stop()
		}
	}
}

func (p *Periodic) UpdatePixels() {
	commits := p.Store.GetDayCommits()

	// Write to the contract from a address with ETH
	
}

func (p *Periodic) MintAndStartAuction() error {
	pixels, err := r.Trixels.GetPixels()
	if err != nil {
		return err
	}

	// Construct image from pixels
	dimension := 200
	upLeft := image.Point{0, 0}
	lowRight := image.Point{dimension, dimension}
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})
	for i := 0; i < dimension; i++ {
		for j := 0; j < dimension; j++ {
			color := ParseHexColor(pixels[i][j])
			img.Set(x, y, color)
		}
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
		Name: time.Now().Format(),
		Image: url,
		Description: "Trixels NFT"
	}

	metaPath := "metadata.json"
	file, err := json.MarshalIndent(data, "", " ")
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
	err = os.Remove(path)
	if err != nil {
		return err
	}


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