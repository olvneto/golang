package main

import "fmt"

func compras(trab1, trab2 bool) (bool, bool, bool) {
	tv50 := trab1 && trab2
	tv32 := trab1 != trab2 // ou exclusivo
	sorvete := trab1 || trab2

	return tv50, tv32, sorvete
}

func main() {
	tv50, tv32, sorvete := compras(true, true)
	fmt.Printf("%t %t %t Saudável:%t", tv50, tv32, sorvete, !sorvete)
}
