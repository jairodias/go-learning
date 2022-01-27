package main

import "fmt"

func somar(numberOne int8, numberTwo int8) int8 {
	return numberOne + numberTwo
}

func calculosMatematicos(numberOne, numberTwo int8) (int8, int8) {
	soma := numberOne + numberTwo
	subtracao := numberOne - numberTwo

	return soma, subtracao
}

func main() {
	soma := somar(10, 15)
	fmt.Println(soma)

	var f = func() {
		fmt.Println("Função f")
	}

	f()

	resultadoSoma, _ := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma)
}
