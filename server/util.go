package server

import (
	"crypto/sha256"
    "image"
	"image/color"
	"fmt"
	"strings"
	"strconv"
)

// ComputePixelHash takes in pixel coordinates and computes the pixel hash.
// It returns the Sha-256 hash of the coordinates.
func ComputePixelHash(x, y uint16) string {
	hashString := fmt.Sprintf("%d,%d", x, y)
	return Sha256(hashString)
}

// Sha256 hashes a payload with Sha-256.
// It returns the hashed string.
func Sha256(payload string) string {
	h := sha256.New()
	h.Write([]byte(payload))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// DecodePixelValues converts the pixel coordinate specifier in the URL into uint coordinates.
// Returns two uint16 representing the coordinates.
func DecodePixelValues(value string) (uint16, uint16) {
	split := strings.Split(value, "-")
	xCoord, _ := strconv.ParseInt(split[0], 0, 16)
	yCoord, _ := strconv.ParseInt(split[1], 0, 16)
	return uint16(xCoord), uint16(yCoord) 
}

// ParseHexColor takes a hex code and converts it to RGBA.
// It returns a color.RGBA struct and an error.
func ParseHexColor(s string) (c color.RGBA, err error) {
    c.A = 0xff
    switch len(s) {
    case 7:
        _, err = fmt.Sscanf(s, "#%02x%02x%02x", &c.R, &c.G, &c.B)
    case 4:
        _, err = fmt.Sscanf(s, "#%1x%1x%1x", &c.R, &c.G, &c.B)
        c.R *= 17
        c.G *= 17
        c.B *= 17
    default:
        err = fmt.Errorf("invalid length, must be 7 or 4")
    }

    return
}

// CreateImage takes pixel hex values and converts them to an image.
// It returns the image.
func CreateImage(pixels *Pixels) image.Image {
	upLeft := image.Point{0, 0}
	lowRight := image.Point{IMAGE_DIMENSIONS, IMAGE_DIMENSIONS}
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})
	for i := 0; i < IMAGE_DIMENSIONS; i++ {
		for j := 0; j < IMAGE_DIMENSIONS; j++ {
			img.Set(i, j, color.White)
		}
	}

	for i := 0; i < len(*pixels); i++ {
		color, err := ParseHexColor((*pixels)[i].Color)
		if err != nil {
			continue
		}
		img.Set(int((*pixels)[i].X), int((*pixels)[i].Y), color)
	}

	return img
}