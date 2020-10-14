package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// Usuario usuarios do banco de dados
type Usuario struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}

//UsuarioHandler analisa o request e delega para a função adequada
func UsuarioHandler(w http.ResponseWriter, r *http.Request) {
	sid := strings.TrimPrefix(r.URL.Path, "/usuarios/")
	id, _ := strconv.Atoi(sid)

	switch {
	case r.Method == "GET" && id > 0:
		usuarioPorID(w, r, id)
	case r.Method == "GET":
		usuarioTodos(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Desculpa... :(")
	}
}

func conexaoBanco() *sql.DB {
	db, err := sql.Open("mysql", "root:password@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func usuarioPorID(w http.ResponseWriter, r *http.Request, id int) {
	db := conexaoBanco()
	defer db.Close()

	var u Usuario
	db.QueryRow("select id, nome from usuarios where id = ?", id).Scan(&u.ID, &u.Nome)

	var usuarios []Usuario
	usuarios = append(usuarios, u)

	publica(w, usuarios)
}

func usuarioTodos(w http.ResponseWriter, r *http.Request) {
	db := conexaoBanco()
	defer db.Close()

	rows, _ := db.Query("select id, nome from usuarios")
	defer rows.Close()

	var usuarios []Usuario
	for rows.Next() {
		var usuario Usuario
		rows.Scan(&usuario.ID, &usuario.Nome)
		usuarios = append(usuarios, usuario)
	}

	publica(w, usuarios)
}

func publica(w http.ResponseWriter, usuarios []Usuario) {
	json, _ := json.Marshal(usuarios)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(json))
}
