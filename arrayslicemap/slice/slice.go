package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3}                               // array
	s1 := []int{1, 2, 3}                                //slice
	fmt.Println(a1, s1)                                 // [1 2 3] [1 2 3]
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1)) // [3]int []int

	//slice não é um array, slice define um pedaço de um array
	a2 := [5]int{1, 2, 3, 4, 5}
	s2 := a2[1:3]
	fmt.Println(a2, s2) //[1 2 3 4 5] [2 3]

	s3 := a2[:2] // novo slice [1 2]
	fmt.Println(a2, s3)

	//slice: tamanho + um ponteiro para um elemento de um array (percorre o tamanho informado)

	s4 := s2[:1]
	fmt.Println(s2, s4)

}
