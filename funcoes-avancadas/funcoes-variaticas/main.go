package main

func soma(numeros ...int) int {
	total := 0

	for _, numero := range numeros {
		total += numero
	}

	return total
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		println(texto, numero)
	}
}

func main() {
	total := soma(1, 2, 3)
	println(total)

	escrever("Olá Mundo", 10, 12, 4)
}
