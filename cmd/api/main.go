package main

import (
	"log"
	"os"

	db "github.com/techerpierre/kasa-api/models"
)

func main() {
	app, err := Setup()

	if err != nil {
		log.Fatalln(err.Error())
	}

	prisma := db.NewClient(
		db.WithDatasourceURL(os.Getenv("DATABASE_URL")),
	)

	if err := prisma.Prisma.Connect(); err != nil {
		panic(err)
	}

	defer func() {
		if err := prisma.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	Instanciate(app, prisma)

	app.Run(os.Getenv("PORT"))
}
