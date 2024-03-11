package belajar_golang_database

import (
	"context"
	"fmt"
	"strconv"
	"testing"
)

func TestTransactionCommit(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	context := context.Background()
	begin, err := db.Begin()
	if err != nil {
		panic(err)
	}
	scriptSql := "insert into comments(email, comment) values (?,?)"

	// do transaction
	for i := 0; i < 10; i++ {
		email := "iqbal" + strconv.Itoa(i) + "@email.com"
		comment := "comment - " + strconv.Itoa(i)
		result, err := begin.ExecContext(context, scriptSql, email, comment)
		if err != nil {
			panic(err)
		}
		insertId, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("Comment id ", insertId)
	}

	err = begin.Commit()
	if err != nil {
		panic(err)
	}
}

func TestTransactionRollback(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	context := context.Background()
	begin, err := db.Begin()
	if err != nil {
		panic(err)
	}
	scriptSql := "insert into comments(email, comment) values (?,?)"

	// do transaction
	for i := 0; i < 10; i++ {
		email := "iqbal" + strconv.Itoa(i) + "@email.com"
		comment := "comment - " + strconv.Itoa(i)
		result, err := begin.ExecContext(context, scriptSql, email, comment)
		if err != nil {
			panic(err)
		}
		insertId, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("Comment id ", insertId)
	}

	err = begin.Rollback()
	if err != nil {
		panic(err)
	}
}
