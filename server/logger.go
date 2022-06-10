package server

import (
	"os"
	"strconv"
	"go.uber.org/zap"
)

var Logger *zap.Logger 

func InitLogger() error {
	prod, err := strconv.ParseBool(os.Getenv("PROD"))
	if err != nil {
		return err 
	}

	if prod {
		Logger, err = zap.NewProduction()
		return err
	}

	Logger, err = zap.NewDevelopment()
	return err
}