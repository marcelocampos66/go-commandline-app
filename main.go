package main

import (
	"commandline-app/app"
	"log"
	"os"
)

func main() {
	application := app.Generate()

	if erro := application.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
