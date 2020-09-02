package main

import "fmt"

type item struct {
	produtoID int
	qtde      float64
	preco     float64
}

type pedido struct {
	userID int
	itens  []item
}

func (p pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.preco * item.qtde
	}
	return total

}

func main() {
	pedido1 := pedido{
		userID: 1,
		itens: []item{
			item{1, 2, 12.10},
			item{2, 1, 23.49},
			item{3, 100, 3.13},
		},
	}

	fmt.Printf("Valor total do pedido: R$ %.2f", pedido1.valorTotal())
}
