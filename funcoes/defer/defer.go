package main

import "fmt"

func obterValorAprovado(valor int) int {
	defer fmt.Println("Fim!")
	if valor > 5000 {
		fmt.Println("Valor alto...")
		return 5000
	}
	fmt.Println("Valor baixo...")
	return valor
}

func main() {
	fmt.Println(obterValorAprovado(6000))
	fmt.Println(obterValorAprovado(3000))
}
