package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5}

	for i, numero := range numeros { //retorna o índice do array (i) e o conteúdo (numero)
		fmt.Printf("%d) %d\n", i+1, numero)
	}

	for _, num := range numeros {
		//o índice não precisa ser usado. Se não colocar o _, o num receberá o índice
		fmt.Println(num)
	}

}
