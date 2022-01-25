package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa    // "HERANÇA"
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança só que não")

	p1 := pessoa{"Jairo", "Dias", 22, 181}
	fmt.Println(p1)

	e1 := estudante{p1, "Análise e Desenvolvimento de Sistemas", "Cruzeiro do Sul"}
	fmt.Println(e1)
}
