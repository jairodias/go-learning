package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)

	go escrever("Olá mundo!", canal)

	println("Depois da função escrever começar a ser executada")

	for mensagem := range canal {
		fmt.Println(mensagem)
	}
}

func escrever(text string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- text // atribuindo o valor
		time.Sleep(time.Second)
	}

	close(canal)
}
