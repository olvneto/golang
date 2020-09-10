package main

import "fmt"

type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

//interfaces são implementadas implicitamente

func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (pr produto) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", pr.nome, pr.preco)
}

func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}

func main() {
	var coisa imprimivel = pessoa{"Ricardo", "Silva"}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	coisa = produto{"Calça Jeans", 79.00}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	p2 := produto{"Calça moletom", 29.00}
	imprimir(p2)

}
