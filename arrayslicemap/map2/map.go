package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"José Joao":      11325.0,
		"Gabriela Silva": 15456.78,
		"Pedro Júnior":   1200.0, //necessário ter a vírgula
	}

	funcsESalarios["Rafael Filho"] = 1350.0

	delete(funcsESalarios, "Inexistente") //não dá erro, simplesmente não faz nada

	for nome, salarios := range funcsESalarios {
		fmt.Println(nome, salarios)
	}

}
