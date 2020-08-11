package main

import "fmt"

func main() {
	//Go não tem aritmética de ponteiros
	i := 1

	var p *int = nil

	p = &i // p aponta para o endereço da variável i
	*p++   // incrementa o valor de i, desreferenciando (pegando o valor)
	i++    // incrementa novamente o valor de i

	fmt.Println(p, *p, i, &i) //0xc00009c090 3 3 0xc00009c090

}
