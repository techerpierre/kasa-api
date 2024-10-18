package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Setup() (*gin.Engine, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	return gin.Default(), nil
}
