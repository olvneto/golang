package main

import "fmt"

func trocar(p1, p2 int) (primeiro int, segundo int) {
	primeiro = p2
	segundo = p1
	return // retorno limpo, já definido no retorno
	//return primeiro, segundo -> também é possível, mas desnecessário
}

func main() {
	a, b := trocar(8, 78)
	fmt.Println(a, b)
}
