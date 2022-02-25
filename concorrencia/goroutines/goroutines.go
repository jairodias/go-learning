package main

import "time"

func main() {
	go escrever("Olá mundo") // goroutine
	escrever("Promando em go!")
}

func escrever(text string) {
	for {
		println(text)
		time.Sleep(time.Second)
	}
}
