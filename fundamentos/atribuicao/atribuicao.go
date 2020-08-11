package main

import "fmt"

func main() {
	var b byte = 3
	fmt.Println(b)

	i := 3
	i += 3
	i -= 3
	i /= 3
	i *= 3
	i %= 3

	fmt.Println(i)

	x, y := 1, 2 // atribuição dupla
	fmt.Println(x, y)

	x, y = y, x //troca de valores
	fmt.Println(x, y)
}
