package main

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	println("Funções Recursivas")

	posicao := uint(10)

	for i := uint(0); i <= posicao; i++ {
		println(fibonacci(i))
	}
}
