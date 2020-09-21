package main

import "fmt"

func rotina(ch chan int) {
	ch <- 1
	fmt.Println("1!")
	ch <- 2
	fmt.Println("2!")
	ch <- 3
	fmt.Println("3!")
	ch <- 4
	fmt.Println("4!")
	ch <- 5
	fmt.Println("5!")
	ch <- 6
}

func main() {
	ch := make(chan int, 3)

	go rotina(ch)

	fmt.Println(<-ch)
}
