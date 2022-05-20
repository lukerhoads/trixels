package server

import (
  	"gorm.io/gorm"
)

type Pixel struct {
	ID uint `gorm:"primaryKey"`
	Hash string `json:"hash" gorm:"type:varchar(64);"`
	X uint16 `json:"x"`
	Y uint16 `json:"y"`
	Color string `json:"color" gorm:"type:varchar(8);"`
	LastAddress string `json:"last_address" gorm:"type:varchar(64);"`
}

type Pixels []Pixel

func NewPixel(x, y uint16) *Pixel {
	return &Pixel{
		Hash: ComputePixelHash(x, y),
		X: x,
		Y: y,
		Color: "#000000",
		LastAddress: "0x0000000000000000000000000000000000000000",
	}
}

func (p *Pixel) GetPixel(db *gorm.DB) {
	if p.Hash == "" {
		p.Hash = ComputePixelHash(p.X, p.Y)
	}

	db.First(&p, "hash = ?", p.Hash)
}

func (p *Pixels) GetPixels(db *gorm.DB) {
	db.Find(&p)
}

func (p *Pixel) UpdatePixel(db *gorm.DB) {
	pixelHash := ComputePixelHash(p.X, p.Y)
	destPixel := Pixel{}
	db.First(&destPixel, "hash = ?", pixelHash)
	if destPixel.validPixel() {
		db.Model(Pixel{}).Where("hash = ?", pixelHash).Updates(p)
	} else {
		newPixel := NewPixel(p.X, p.Y)
		db.Create(newPixel)
	}
}

func (p *Pixel) validPixel() bool {
	return p.Color != ""
}

type Metadata struct {
	Name string `json:"name"`
	Image string `json:"image"`
	Description string `json:"description"`
}