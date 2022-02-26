package enderecos

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {
	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Rua"},
		{"Praça das Rosas", "Tipo invàlido"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		{"ESTRADA QUALQUER", "Estrada"},
		{"", "Tipo inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s",
				cenario.retornoEsperado,
				tipoDeEnderecoRecebido)
		}
	}

}
