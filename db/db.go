package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

/* Conectar com a base de dados */
func Conectar() *sql.DB {
	conexao := "user=postgres dbname=crud_go password=celta sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
