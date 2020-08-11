package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println(a+b, a-b, a/b, a*b, a%b) //5 1 1 6 1

	//bitwise
	fmt.Println(a&b, a|b, a^b) //10 11 01

	c := 3.0
	d := 2.0

	fmt.Println(math.Max(float64(a), float64(b)), math.Min(c, d), math.Pow(c, d)) //3 2 9

}
