package server

import (
	"time"
	"gorm.io/gorm"
	"errors"
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
	TokenID     uint `json:"tokenID" gorm:"primaryKey"`
	MetadataUrl string `json:"metadataUrl"`
	CreatedAt *time.Time `json:"createdAt"`
	live bool
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

func (t *Trixel) GetTrixel(db *gorm.DB) bool {
	return db.First(&t, "token_id = ?", t.TokenID).Error != gorm.ErrRecordNotFound
}

func (t *Trixel) GetLiveTrixel(db *gorm.DB) bool {
	return db.First(&t, "token_id = ?", t.TokenID).Error != gorm.ErrRecordNotFound
}

func (t *Trixel) AddLiveTrixel(db *gorm.DB) error {
	var trixel []Trixel
	var count int64
	db.First(&trixel, "live = ?", true).Count(&count)
	if count > 1 {
		return errors.New("Cannot add another live auction")
	}

	db.Create(t)
	return nil
}

func (t *Trixel) DeleteLiveTrixel(db *gorm.DB) {
	db.Exec("DELETE FROM trixels WHERE live = ?", true)
}

func (t *Trixel) AddTrixel(db *gorm.DB) {
	db.Create(t)
}

func (t *Trixels) GetTrixels(db *gorm.DB) {
	db.Find(&t).Where("live = ?", false)
}

type ServerError struct {
	Message string `json:"message"`
}