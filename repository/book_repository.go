package repository

import (
	"context"
	"crud/model"
)

type BookRespository interface{
	Save(ctx context.Context , book model.Book)
	Update(ctx context.Context , book model.Book)
	Delete(ctx context.Context , bookId int)
	FindById(ctx context.Context , bookId int)(model.Book , error)
	FindAll(ctx context.Context )[]model.Book
}