package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	//cuidado...
	fmt.Println("Teste " + string(123)) //Teste {

	//int para string
	fmt.Println("Teste " + strconv.Itoa(123)) // Teste 123

	//string para int
	num, _ := strconv.Atoi("123") // retorna dois valores, o número esperado e o erro (_), caso a
	// string não for um valor inteiro
	fmt.Println(num - 122) // 1

	b, _ := strconv.ParseBool("true") // ou ParseBool("1")
	if b {
		fmt.Println("Verdadeiro")
	}

}
