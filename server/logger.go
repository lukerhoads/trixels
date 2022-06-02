package server

import (
	"os"
	"strconv"
	"go.uber.org/zap"
)

var Logger *zap.Logger 

func InitLogger() (*zap.Logger, error) {
	prod, err := strconv.ParseBool(os.Getenv("PROD"))
	if err != nil {
		return nil, err 
	}

	if prod {
		return zap.NewProduction()
	}

	return zap.NewDevelopment()
}