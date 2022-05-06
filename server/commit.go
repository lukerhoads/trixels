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
	X uint16 
	Y uint16 
	Color string 
	Address string
}