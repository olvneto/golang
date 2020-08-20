package main

import "fmt"

func main() {
	//maps devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[1234] = "Maria"
	aprovados[4567] = "Pedro"
	aprovados[1245] = "Ana"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d) ", nome, cpf)
	}
	fmt.Println()
	delete(aprovados, 1245)

	fmt.Println(aprovados)
}
