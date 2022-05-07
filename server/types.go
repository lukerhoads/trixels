package main

import (
	"time"
  	"gorm.io/gorm"
)

type Commit struct {
	CreatedAt time.Time `json:"-"`
	X uint16 `json:"x"`
	Y uint16 `json:"y"`
	Color string `json:"color"`
	Address string `json:"address"`
}

func (c *Commit) CreateCommit(db *gorm.DB) error {
	result := db.Create(&c)
	if (result.Error != nil) {
		return result.Error
	}

	return nil
}

func (c *Commit) GetDayCommits(db *gorm.DB) []Commit {
	var commits []Commit
	year, month, day := time.Now().Date()
    today := time.Date(year, month, day, 0, 0, 0, 0, time.Now().Location())
	db.Where("created_at >", today).Find(&commits)
	return commits
}

type Pixel struct {
	ID uint `json:"id"`
	PixelID uint16 `json:"pixel_id"` 
	X uint16 `json:"x"`
	Y uint16 `json:"y"`
	Color string `json:"color" gorm:"type:varchar(8);"`
	LastAddress string `json:"last_address" gorm:"type:varchar(64);"`
}

func (p *Pixel) UpdatePixel(db *gorm.DB) {
	computedId := computeId(p.X, p.Y)
	p.PixelID = computedId
	destPixel := Pixel{}
	db.First(&destPixel, "pixel_id = ?", computedId)
	if destPixel.validPixel() {
		db.Model(Pixel{}).Where("pixel_id = ?", computedId).Updates(p)
	} else {
		db.Create(p)
	}
}

// 0, 0 -> 0
// 1, 20 -> 220
func computeId(x uint16, y uint16) uint16 {
	return (y * 200) + x
}

func (p *Pixel) validPixel() bool {
	return p.Color != ""
}

type DimensionedPixel struct {
	X string `json:"x"`
	Y string `json:"y"`
	Color string `json:"color"`
}

type Metadata struct {
	Name string `json:"name"`
	Image string `json:"image"`
	Description string `json:"description"`
}