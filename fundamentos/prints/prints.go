package main

import "fmt"

func main() {
	fmt.Print("Imprime tudo na mesma linha.")
	fmt.Println() //vai para a linha de baixo
	fmt.Println("Imprime com quebra de linha no final")

	x := 3.141516

	fmt.Println("O valor de x é", x)

	xs := fmt.Sprint(x) // O valor de x é 3.141516 (transforma para string e joga em xs)

	fmt.Println("O valor de x é " + xs) //O valor de x é 3.141516

	fmt.Printf("O valor de x é %.2f.", x) //O valor de x é 3.14.

	a := 1
	b := 1.9999
	c := false
	d := "opa"

	fmt.Printf("\n %d %f %.1f %t %s", a, b, b, c, d) //1 1.999900 2.0 false opa
	fmt.Printf("\n %v %v %v %v", a, b, c, d)         //1 1.9999 false opa
}
