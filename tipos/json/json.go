package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	//struct par json
	p := produto{1, "Notebook", 1990.0, []string{"Promoção", "Eletrônicos"}}
	p1Json, _ := json.Marshal(p)
	fmt.Println(string(p1Json))

	//json para struct
	var p2 produto
	jsonString := `{"id":1, "nome":"Caneta", "preco":2.90,"tags":["Papelaria", "Importado"]}`
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2)
}
