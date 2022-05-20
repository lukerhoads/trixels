package server

import (
	"crypto/sha256"
	"fmt"
	"strings"
	"strconv"
    "image/color"
)

func ComputePixelHash(x, y uint16) string {
	hashString := fmt.Sprintf("%d,%d", x, y)
	return Sha256(hashString)
}

func Sha256(payload string) string {
	h := sha256.New()
	h.Write([]byte(payload))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func DecodePixelValues(value string) (uint16, uint16) {
	split := strings.Split(value, "-")
	xCoord, _ := strconv.ParseInt(split[0], 0, 16)
	yCoord, _ := strconv.ParseInt(split[1], 0, 16)
	return uint16(xCoord), uint16(yCoord) 
}

func ParseHexColor(s string) (c color.RGBA, err error) {
    c.A = 0xff
    switch len(s) {
    case 7:
        _, err = fmt.Sscanf(s, "#%02x%02x%02x", &c.R, &c.G, &c.B)
    case 4:
        _, err = fmt.Sscanf(s, "#%1x%1x%1x", &c.R, &c.G, &c.B)
        // Double the hex digits:
        c.R *= 17
        c.G *= 17
        c.B *= 17
    default:
        err = fmt.Errorf("invalid length, must be 7 or 4")
    }

    return
}