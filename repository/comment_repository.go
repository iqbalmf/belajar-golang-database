package repository

import (
	"belajar-golang-database/entity"
	"context"
	"database/sql"
	"errors"
	"strconv"
)

type CommentRepository interface {
	Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error)
	FindAll(ctx context.Context) ([]entity.Comment, error)
	FindById(ctx context.Context, id int32) (entity.Comment, error)
	UpdateById(ctx context.Context, id int32, comment entity.Comment) (entity.Comment, error)
	DeleteById(ctx context.Context, id int32) (string, error)
	DeleteAll(ctx context.Context) (string, error)
}

type commentRepositoryImpl struct {
	Db *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryImpl{Db: db}
}

func (c *commentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	script := "insert into comments(email, comment) values (?,?)"
	result, err := c.Db.ExecContext(ctx, script, comment.Email, comment.Comment)
	if err != nil {
		return comment, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}
	comment.Id = int32(id)
	return comment, nil
}

func (c *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	script := "select id, email, comment from comments"
	rows, err := c.Db.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var comments []entity.Comment
	for rows.Next() {
		comment := entity.Comment{}
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		comments = append(comments, comment)
	}
	return comments, nil
}

func (c *commentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	script := "select id, email, comment from comments where id = ? limit 1"
	rows, err := c.Db.QueryContext(ctx, script, id)
	comment := entity.Comment{}
	if err != nil {
		return comment, err
	}
	defer rows.Close()
	if rows.Next() {
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		return comment, nil
	} else {
		return comment, errors.New("Id " + strconv.Itoa(int(id)) + " not found")
	}
}

func (c *commentRepositoryImpl) UpdateById(ctx context.Context, id int32, comment entity.Comment) (entity.Comment, error) {
	script := "update comments set email = ? where id = ?"
	_, err := c.Db.ExecContext(ctx, script, comment.Email, id)
	updateComment := entity.Comment{
		Id: id, Email: comment.Email,
	}
	if err != nil {
		return comment, err
	}
	return updateComment, nil
}

func (c *commentRepositoryImpl) DeleteById(ctx context.Context, id int32) (string, error) {
	script := "delete from comments where id = ?"
	_, err := c.Db.ExecContext(ctx, script, id)

	if err != nil {
		return "", err
	}
	return "Delete Success", nil
}

func (c *commentRepositoryImpl) DeleteAll(ctx context.Context) (string, error) {
	script := "delete from comments"
	_, err := c.Db.ExecContext(ctx, script)

	if err != nil {
		return "", err
	}
	return "Delete All Success", nil
}
