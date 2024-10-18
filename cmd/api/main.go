package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	app, err := Setup()

	if err != nil {
		log.Fatalln(err.Error())
	}

	app.GET("/api/doc", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hello"})
	})

	app.Run(os.Getenv("PORT"))
}
