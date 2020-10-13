package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:password@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into usuarios(id, nome) values(?, ?)")
	stmt.Exec(20, "Bia")
	stmt.Exec(21, "Carlos")
	_, err = stmt.Exec(1, "Tiago") //chave duplicada

	if err != nil { // nada é realizado no banco, pois há o rollback
		tx.Rollback()
		log.Fatal(err)
	} // o erro tem que ser tratado antes do commit

	tx.Commit() // não é executado
}
