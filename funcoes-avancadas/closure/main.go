package main

func closure() func() {
	texto := "Dentro da função closure"

	funcao := func() {
		println(texto)
	}

	return funcao
}

func main() {
	println("Função closure")

	texto := "Dentro da função main"
	println(texto)

	funcaoNova := closure()

	funcaoNova()
}
