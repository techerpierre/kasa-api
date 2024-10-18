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

	app.GET("/api/hello-world", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello world !"})
	})

	app.Run(os.Getenv("PORT"))
}
