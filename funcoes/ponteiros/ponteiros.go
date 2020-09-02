package main

import "fmt"

func inc1(n int) {
	n++
}

func inc2(n *int) {
	*n++ //pega o valor de n e incrementa
}

func main() {
	n := 1

	inc1(n)
	fmt.Println(n) //passa apenas o valor de n

	inc2(&n)
	fmt.Println(n) //passa o endereço de n e imprime o valor incrementado
}
