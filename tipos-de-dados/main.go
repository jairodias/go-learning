package main

import (
	"errors"
	"fmt"
)

func main() {
	// números inteiros
	numero := 100000000000
	fmt.Println(numero)

	var numero1 uint16 = 10000 // não aceita numero negativos
	fmt.Println(numero1)

	// alias
	// INT32 = RUNE
	var numero2 rune = 123456
	fmt.Println(numero2)

	// BYTE = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)

	// FLOAT32
	var numero5 float32 = 123.45
	fmt.Println(numero5)

	// FLOAT64
	var numero6 float32 = 1230000.45
	fmt.Println(numero6)

	var str string = "text"
	fmt.Println(str)

	str2 := "Text"
	fmt.Println(str2)

	char := 'A'
	fmt.Println(char)

	var text string
	fmt.Println(text)

	var booleano1 bool = true
	fmt.Println(booleano1)

	var error error = errors.New("Erro interno")
	fmt.Println(error)
}
