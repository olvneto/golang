package main

import "fmt"

func main() {
	s := make([]int, 10)
	s[9] = 12
	fmt.Println(s) // [0 0 0 0 0 0 0 0 0 12]

	s = make([]int, 10, 20)
	fmt.Println(s, len(s), cap(s)) //[0 0 0 0 0 0 0 0 0 0] 10 20

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s, len(s), cap(s)) //[0 0 0 0 0 0 0 0 0 0 1 2 3 4 5 6 7 8 9 0] 20 20

	s = append(s, 1)
	fmt.Println(s, len(s), cap(s)) //[0 0 0 0 0 0 0 0 0 0 1 2 3 4 5 6 7 8 9 0 1] 21 40
}
