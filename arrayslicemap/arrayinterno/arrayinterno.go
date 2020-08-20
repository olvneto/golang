package main

import "fmt"

func main() {
	s := make([]int, 10, 20)
	s2 := append(s, 1, 2, 3)
	fmt.Println(s, s2)

	s[0] = 7
	fmt.Println(s, s2) //[7 0 0 0 0 0 0 0 0 0] [7 0 0 0 0 0 0 0 0 0 1 2 3]
}
