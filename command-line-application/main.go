package main

import (
	"command-line-application/app"
	"log"
	"os"
)

func main() {
	println("Executando a aplicação")

	application := app.Gerar()

	if error := application.Run(os.Args); error != nil {
		log.Fatal(error)
	}
}
