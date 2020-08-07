package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	//números inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(3200))

	// sem sinal (positivos) uint8 uint16 uint32 uint64
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	// com sinal int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1)) //O tipo de i1 é int

	var i2 rune = 'a'                           //representa um mapeamento da tabela Unicode (int32)
	fmt.Println("O rune é", reflect.TypeOf(i2)) //O rune é int32
	fmt.Println(i2)                             //97

	//números reais
	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))         //O tipo de x é float32
	fmt.Println("O tipo de 49.99 é", reflect.TypeOf(49.99)) //O tipo de 49.99 é float64

	//boolean
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo)) //O tipo de bo é bool
	fmt.Println(!bo)                                  //false

	//string
	s := "Olá, meu nome é estudante"
	fmt.Println(s + "!")                         //Olá, meu nome é estudante!
	fmt.Println("O tamanho da string é", len(s)) //O tamanho da string é 27

	//string com múltiplas linhas
	s1 := `Olá
	meu
	nome
	é estudante
	`
	fmt.Println(s1) //Olá
	//meu
	//nome
	//é estudante

	char := 'a'                                           // não existe o tipo char em Go
	fmt.Println("O tipo de char é", reflect.TypeOf(char)) //O tipo de char é int32
	fmt.Println(char)                                     //97

}
