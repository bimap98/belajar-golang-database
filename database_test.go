package belajar_golang_database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestEmpty(t *testing.T) {

}

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:abcd12345@tcp(localhost:3306)/belajar_golang_database")

	if err != nil {
		panic(err)
	}

	defer db.Close()

}
