package main

func recuperarExecucao() {
	if r := recover(); r != nil {
		println("Execução recuperada com sucesso.")
	}
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A MÉDIA É EXATAMENTE 6!")
}

func main() {
	println("Função DEFER")

	println(alunoEstaAprovado(6, 6))
	println("Pós execução")
}
