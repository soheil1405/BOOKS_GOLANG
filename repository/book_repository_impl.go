package repository

import (
	"context"
	"crud/helper"
	"crud/model"
	"database/sql"
	"errors"
)

type BookRespositoryImpl struct {
	Db *sql.DB
}

func NewBookRespository(Db *sql.DB) BookRespository {
	return &BookRespositoryImpl{Db: Db}
}

func (b *BookRespositoryImpl) Save(ctx context.Context, book model.Book) {

	tx, err := b.Db.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	SQL := "insert into book(name) value($1)"

	_, err = tx.ExecContext(ctx, SQL, book.Name)

	helper.PanicIfErr(err)

}

func (b *BookRespositoryImpl) Update(ctx context.Context, book model.Book) {

	tx, err := b.Db.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	SQL := "update book set name=$1 weher id =$2"
	_, err = tx.ExecContext(ctx, SQL, book.Name, book.Id)

	helper.PanicIfErr(err)
}

func (b *BookRespositoryImpl) Delete(ctx context.Context, bookId int) {

	tx, err := b.Db.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	SQL := "delete from book where id=$id"
	_, err = tx.ExecContext(ctx, SQL, bookId)
	helper.PanicIfErr(err)
}

func (b BookRespositoryImpl) FindById(ctx context.Context, bookId int) (model.Book, error) {

	tx, err := b.Db.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	SQL := "select  * from books where id=$id"

	result, err := tx.QueryContext(ctx, SQL)

	helper.PanicIfErr(err)
	defer result.Close()
	book := model.Book{}

	if result.Next() {
		err := result.Scan(&bookId, &book.Name)
		helper.PanicIfErr(err)
		return book, nil
	} else {
		return book, errors.New("book id not found")
	}

}

func (b *BookRespositoryImpl) FindAll(ctx context.Context) []model.Book {

	tx, err := b.Db.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	SQL := "select id,Name from book"

	result, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfErr(err)
	defer result.Close()

	var books []model.Book

	for result.Next() {
		book := model.Book{}
		err := result.Scan(&book.Id, &book.Name)
		helper.PanicIfErr(err)
		books = append(books, book)
	}

	return books
}
