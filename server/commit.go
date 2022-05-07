package main

import (
	"time"
)

type Commit struct {
	CreatedAt time.Time
	X uint16
	Y uint16 
	Color string 
	Address string
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