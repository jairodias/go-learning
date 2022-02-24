package main

func inverterSinal(numero int) int {
	return numero * -1
}

func inverterSinalComPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {
	println("Função PONTEIROS")
	numero := 20
	numeroInvertido := inverterSinal(numero)
	println(numeroInvertido)
	println(numero)

	novoNumero := 40
	println(novoNumero)
	inverterSinalComPonteiro(&novoNumero)
	println(novoNumero)
}
