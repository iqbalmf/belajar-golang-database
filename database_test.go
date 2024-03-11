package belajar_golang_database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "admin:passwordadmin132@tcp(8.215.0.67:3306)/belajar_golang_database")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
