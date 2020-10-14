package main

import (
	"log"
	"net/http"
)

//roda a partir do terminal com: $ go run static.go
func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
