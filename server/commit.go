package main

import (
	"time"
)

type Commit struct {
	CreatedAt time.Time `json:"-"`
	X uint16 `json:"x"`
	Y uint16 `json:"y"`
	Color string `json:"color"`
	Address string `json:"address"`
}

type Pixel struct {
	Color string 
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