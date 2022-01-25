package main

import "fmt"

type endereco struct {
	logradouro string
	numero     uint8
}

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "Jairo Dias"
	u.idade = 22
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua dos Bobos", 0}

	usuario2 := usuario{"Jairo Dias", 23, enderecoExemplo}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Jairo Dias"}
	fmt.Println(usuario3)
}
