package main

import (
	"time"
)

func main() {
	i := 0

	for i < 2 {
		i++
		println("Incrementando i", i)
		time.Sleep(time.Second)
	}

	for j := 1; j < 2; j++ {
		println("Incrementando j", j)
		time.Sleep(time.Second)
	}

	nomes := [3]string{"João", "Davi", "Lucas"}

	for _, nome := range nomes {
		println(nome)
	}

	for indice, letra := range "PALAVRA" {
		println(indice, letra)
	}

	usuario := map[string]string{
		"primeiro": "João",
		"ultimo":   "Silva",
	}

	for chave, valor := range usuario {
		println(chave, valor)
	}
}
