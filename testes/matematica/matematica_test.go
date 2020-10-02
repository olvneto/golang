package matematica

import "testing"

const erroPadrao = "Valor esperado %v, mas o resultado encontrado foi %v."

func TestMedia(t *testing.T) { // assinatura padrão de um método de teste
	t.Parallel() // deve começar com Test... e ter os parâmetros t *testing.T
	valorEsperado := 7.28
	valor := Media(7.2, 9.9, 6.1, 5.9)

	if valor != valorEsperado {
		t.Errorf(erroPadrao, valorEsperado, valor)
	}
} // pode rodar o teste no terminal, quando na pasta de
// testes, com: go test ./...
