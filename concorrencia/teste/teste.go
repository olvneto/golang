package main

import (
	"fmt"
	"time"
)

func isPrimo(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func primos(n int, c chan int, p chan int) {
	inicio := 2
	for i := 0; i < n; i++ {
		for primo := inicio; ; primo++ {
			if isPrimo(primo) {
				c <- primo
				inicio = primo + 1
				p <- i
				time.Sleep(time.Millisecond * 180)
				break
			}
		}
	}
	close(c)
}

func main() {
	c := make(chan int, 30)
	p := make(chan int)
	go primos(40, c, p)
	for primo := range c {
		fmt.Printf("%d ", primo)
		if <-p == 20 {
			fmt.Println()
		}
	}
	fmt.Println("Fim!")
}
