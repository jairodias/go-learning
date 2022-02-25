package main

import (
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func() {
		escrever("Olá mundo!")
		waitGroup.Done()
	}()

	go func() {
		escrever("Promando em go!")
		waitGroup.Done()
	}()

	waitGroup.Wait()
}

func escrever(text string) {
	for i := 0; i < 5; i++ {
		println(text)
		time.Sleep(time.Second)
	}
}
