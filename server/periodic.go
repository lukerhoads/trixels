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
	*TrixelsAuctionHouse
	*ethclient.Client
	dayTicker *time.Ticker 
	twoWeekTicker *time.Ticker 
	devChan chan string
	quit chan struct{}
	privateKey *ecdsa.PrivateKey
}

func NewPeriodic(store *Store, privateKey *ecdsa.PrivateKey, client *ethclient.Client, trixels *Trixels, trixelsAuctionHouse *TrixelsAuctionHouse, dayTicker *time.Ticker, twoWeekTicker *time.Ticker, devChan chan string, quit chan struct{}) *Periodic {
	return &Periodic{
		Store: store,
		Trixels: trixels,
		TrixelsAuctionHouse: trixelsAuctionHouse,
		Client: client,
		dayTicker,
		twoWeekTicker,
		devChan,
		quit,
		privateKey,
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
		// Simply for demo purposes
		case v := <-p.devChan:
			if v == "mint" {
				p.MintAndStartAuction()
			} else if v == "updatepixels" {
				p.UpdatePixels()
			}
		case <-quit:
			p.dayTicker.Stop()
			p.twoWeekTicker.Stop()
		}
	}
}

func (p *Periodic) StartUpdater() {
	// Setup ethereum event log/listener for an update transaction
	contractAddress := common.HexToAddress(address)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
		Topics: {{common.HexToHash("0xdfda710484c7dc3d855d2d80e0f7832b0f90247a4a297036e424961e0c590bee")}}
	}

	logs := make(chan types.Log)
	sub, err := p.Client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		return err
	}

	for {
		select {
		case err := <-sub.Err():
		case vLog := <-logs:
			// create a commit based off the transaction log
			log.Println(vLog)
		}
	}
}

func (p *Periodic) UpdatePixels() {
	commits := p.Store.GetDayCommits()

	// Write to the contract from a address with ETH
	keyedTransactor := p.GenKeyedTransactor()

	// Convert commits to new pixels
	for _, c := range commits {

	}

	tx, err := p.Trixels.MassChangePixels()
	if err != nil {
		return err 
	}

	log.Println(tx.Hash().Hex())
	return nil
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

	skyNetId := metaUrl[strings.LastIndex(metaUrl, "/")+1:]
	log.Println(skyNetId)
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
	nonce, err := p.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := p.client.SuggestGasPrice(context.Background())
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