package belajar_golang_database

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	t.Skip("need to clear db first")
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	scriptSql := "INSERT into customer(id, name) values ('iqbal2', 'Iqbal')"
	_, err := db.ExecContext(ctx, scriptSql)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success Insert new Customer")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	scriptSql := "select id, name from customer"
	rows, err := db.QueryContext(ctx, scriptSql)
	if err != nil {
		panic(err)
	}

	for rows.Next() { //melakukan iterasi untuk show result query, selama next() true data masih ada
		var id, name string
		err := rows.Scan(&id, &name) //urutan sesuai dengan query
		if err != nil {
			panic(err)
		}
		fmt.Println("id:", id)
		fmt.Println("name:", name)
	}
	defer rows.Close()
}

func TestQuerySqlComplex(t *testing.T) {
	t.Skip("need to clear db first")
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	scriptSql := "select id, name, email, balance, rating, birth_date, married, created_at from customer"
	rows, err := db.QueryContext(ctx, scriptSql)
	if err != nil {
		panic(err)
	}

	for rows.Next() { //melakukan iterasi untuk show result query, selama next() true data masih ada
		var id, name, email string
		var balance int
		var rating float32
		var birthDate, createdAt time.Time //adding 'parseTime=true' di config open db
		var married bool
		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt) //urutan sesuai dengan query
		if err != nil {
			panic(err)
		}
		fmt.Println("id:", id, "name", name, "email", email, "balance", balance, "rating", rating, "birth-date", birthDate, "married", married, "created at", createdAt)
	}
	defer rows.Close()
}

func TestQuerySqlNullTipeData(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	scriptSql := "select id, name, email, balance, rating, birth_date, married, created_at from customer"
	rows, err := db.QueryContext(ctx, scriptSql)
	if err != nil {
		panic(err)
	}

	for rows.Next() { //melakukan iterasi untuk show result query, selama next() true data masih ada
		var id, name string
		var email sql.NullString
		var balance sql.NullInt32
		var rating sql.NullFloat64
		var birthDate, createdAt sql.NullTime //adding 'parseTime=true' di config open db
		var married sql.NullBool
		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt) //urutan sesuai dengan query
		if err != nil {
			panic(err)
		}
		fmt.Println("id:", id, "name", name, "email", email, "balance", balance, "rating", rating, "birth-date", birthDate, "married", married, "created at", createdAt)
		if email.Valid {
			fmt.Println("email", email.String)
		}
	}
	defer rows.Close()
}

func TestSqlInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin"
	password := "admin"

	scriptSql := "select username from user where username = ? and password = ? limit 1"
	rows, err := db.QueryContext(ctx, scriptSql, username, password)
	if err != nil {
		panic(err)
	}
	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses Login", username)
	} else {
		fmt.Println("Gagal Login")
	}
	defer rows.Close()
}

func TestExecInjectionSafe(t *testing.T) {
	t.Skip("need to clear db first")
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin2"
	password := "admin2"

	scriptSql := "insert into user(username, password) values (?, ?)"
	_, err := db.ExecContext(ctx, scriptSql, username, password)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success Insert new User", username)
}

func TestCommentAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "test@test.com"
	comment := "test comment"

	scriptSql := "insert into comments(email, comment) values (?, ?)"
	insertId, err := db.ExecContext(ctx, scriptSql, email, comment)
	if err != nil {
		panic(err)
	}
	id, err := insertId.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println("Success Insert new Comments", id)
}

/*penggunaan untuk query secara berulang dengan data param yang sama*/
func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	scriptSql := "insert into comments(email, comment) values (?,?)"
	prepareContext, err := db.PrepareContext(ctx, scriptSql) // disini
	if err != nil {
		panic(err)
	}
	defer prepareContext.Close() //jangan lupa diclose

	for i := 0; i < 10; i++ {
		email := "iqbal" + strconv.Itoa(i) + "@email.com"
		comment := "comment - " + strconv.Itoa(i)
		result, err := prepareContext.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}
		insertId, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("Comment id ", insertId)
	}
}
