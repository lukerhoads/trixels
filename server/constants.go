package server

import "time"

const (
	IMAGE_PATH       = "image.png"
	IMAGE_DIMENSIONS = 30

	METADATA_PATH = "metadata.json"
)

var (
	PixelUpdateTime = 5 * time.Minute
)
