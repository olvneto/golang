package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	//fale("Maria", "Por que você não fala?", 3)
	//fale("João", "Só posso falar depois de você", 1)

	//go fale("Maria", "Ei...", 500)//go routine
	//go fale("João", "Opa...", 500)// não imprime nada pois o programa termina antes

	go fale("Maria", "Entendi!!", 10)
	fale("João", "Parabéns", 5) //quando termina, o programa é encerrado
	// independente da go routine anterior
}
