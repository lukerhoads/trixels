package server

import (
	"crypto/sha256"
	"fmt"
	"strings"
	"strconv"
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