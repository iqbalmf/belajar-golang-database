package repository

import (
	belajar_golang_database "belajar-golang-database"
	"belajar-golang-database/entity"
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestCommentInsert(t *testing.T) {
	var db = belajar_golang_database.GetConnection()
	defer db.Close()
	commentRepository := NewCommentRepository(db)

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "test@test.com",
		Comment: "test comment repository",
	}

	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestCommentRepositoryImpl_FindById(t *testing.T) {
	var db = belajar_golang_database.GetConnection()
	defer db.Close()
	commentRepository := NewCommentRepository(db)

	result, err := commentRepository.FindById(context.Background(), 67)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestCommentRepositoryImpl_FindAll(t *testing.T) {
	var db = belajar_golang_database.GetConnection()
	defer db.Close()
	commentRepository := NewCommentRepository(db)

	result, err := commentRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}
	for _, comment := range result {
		fmt.Println(comment)
	}
}

func TestCommentRepositoryImpl_UpdateById(t *testing.T) {
	var db = belajar_golang_database.GetConnection()
	defer db.Close()
	commentRepository := NewCommentRepository(db)
	comment := entity.Comment{
		Id:    66,
		Email: "testemailbaru@test.com",
	}
	result, err := commentRepository.UpdateById(context.Background(), 67, comment)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestCommentRepositoryImpl_DeleteById(t *testing.T) {
	var db = belajar_golang_database.GetConnection()
	defer db.Close()
	commentRepository := NewCommentRepository(db)
	result, err := commentRepository.DeleteById(context.Background(), 67)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestDeleteAll(t *testing.T) {
	var db = belajar_golang_database.GetConnection()
	defer db.Close()
	commentRepository := NewCommentRepository(db)
	result, err := commentRepository.DeleteAll(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
