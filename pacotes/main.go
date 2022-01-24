package main

import (
	"fmt"
	"module/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo no arquivo main")
	auxiliar.Escrever()

	error := checkmail.ValidateFormat("123")
	fmt.Println(error)
}
