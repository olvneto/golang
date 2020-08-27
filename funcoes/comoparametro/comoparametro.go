package main

import "fmt"

func multiplicacao(p1, p2 int) int {
	return p1 * p2
}

func exec(funcao func(int, int) int, p1, p2 int) int {
	return funcao(p1, p2)
}

func main() {
	fmt.Println(exec(multiplicacao, 4, 5))
}
