package server

import (
	"regexp"
	"time"
)

const (
	IMAGE_PATH       = "image.png"
	IMAGE_DIMENSIONS = 30

	METADATA_PATH = "metadata.json"
)

var (
	PixelUpdateTime = 5 * time.Minute

	HexVerificationRegexp = regexp.MustCompile("^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$")
)
