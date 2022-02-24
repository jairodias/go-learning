package main

func alunoEstaAprovado(n1, n2 float32) bool {
	defer println("Média calculada. Resultado será retornado")
	println("Entrando na função para verificar se o aluno está aprovado.")

	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}

	return false
}

func main() {
	println("Função DEFER")

	println(alunoEstaAprovado(7, 8))
}
