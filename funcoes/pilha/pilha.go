package main

import "runtime/debug"

func f1() {
	debug.PrintStack()
}

func f2() {
	f1()
}

func f3() {
	f2()
}

func main() {
	f3()
}
