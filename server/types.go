package server

import (
	"log"
	"time"
	"gorm.io/gorm"
)

// Pixel represents one pixel on the canvas.
type Pixel struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	UpdatedAt 	*time.Time `json:"updatedAt"`
	Hash        string `json:"hash" gorm:"type:varchar(64);"`
	X           uint16 `json:"x"`
	Y           uint16 `json:"y"`
	Color       string `json:"color" gorm:"type:varchar(8);"`
	Editor string `json:"editor" gorm:"type:varchar(64);"`
}

type Pixels []Pixel

func NewPixel(x, y uint16) *Pixel {
	return &Pixel{
		Hash:        ComputePixelHash(x, y),
		X:           x,
		Y:           y,
		Color:       "#000000",
		Editor: "",
	}
}

func (p *Pixel) GetPixel(db *gorm.DB) bool {
	if p.Hash == "" {
		p.Hash = ComputePixelHash(p.X, p.Y)
		log.Println("Computed hash: ", p.Hash)
	}

	return db.First(&p, "hash = ?", p.Hash).Error != gorm.ErrRecordNotFound
}

func (p *Pixels) GetPixels(db *gorm.DB) {
	db.Find(&p)
}

func (p *Pixel) CreateDefaultPixel(db *gorm.DB) error {
	return db.Create(NewPixel(p.X, p.Y)).Error
}

func (p *Pixel) CreatePixel(db *gorm.DB) error {
	p.Hash = ComputePixelHash(p.X, p.Y)
	return db.Create(p).Error
}

// func (p *Pixel) UpdatePixel(db *gorm.DB) {
// 	pixelHash := ComputePixelHash(p.X, p.Y)
// 	destPixel := Pixel{}
// 	db.First(&destPixel, "hash = ?", pixelHash)
// 	if destPixel.validPixel() {
// 		db.Model(Pixel{}).Where("hash = ?", pixelHash).Updates(p)
// 	} else {
// 		log.Println("creating")
// 		newPixel := NewPixel(p.X, p.Y)
// 		db.Create(newPixel)
// 	}
// }

func (p *Pixel) UpdateExistingPixel(db *gorm.DB, valid bool) {
	pixelHash := ComputePixelHash(p.X, p.Y)
	if valid {
		db.Model(Pixel{}).Where("hash = ?", pixelHash).Updates(p)
	} else {
		newPixel := NewPixel(p.X, p.Y)
		db.Create(newPixel)
	}
}

func (p *Pixel) validPixel() bool {
	return p.Color != ""
}

// Metadata represents a trixel's metadata
type Metadata struct {
	Name        string `json:"name"`
	Image       string `json:"image"`
	Description string `json:"description"`
}

func NewMetadata(name string, image string, description string) *Metadata {
	return &Metadata{
		Name:     name,
		Image: image,
		Description: description,
	}
}

// Trixel represents a minted Trixel, which enables redirects, which can ultimately support on-chain mints.
type Trixel struct {
	TokenID     uint `json:"token_id" gorm:"primaryKey"`
	MetadataUrl string `json:"metadata_url"`
}

type Trixels []Trixel

func NewTrixel(tokenID uint, metadataUrl string) *Trixel {
	return &Trixel{
		TokenID:     tokenID,
		MetadataUrl: metadataUrl,
	}
}

func (t *Trixel) GetTrixelCount(db *gorm.DB) (count int64) {
	db.Model(&Trixel{}).Count(&count)
	return
}

func (t *Trixel) GetTrixel(db *gorm.DB) {
	db.First(&t, "token_id = ?", t.TokenID)
}

func (t *Trixel) AddTrixel(db *gorm.DB) {
	db.Create(t)
}

func (t *Trixels) GetTrixels(db *gorm.DB) {
	db.Find(&t)
}

type ServerError struct {
	Message string `json:"message"`
}