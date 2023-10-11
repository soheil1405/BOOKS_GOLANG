package Service

import (
	"context"
	"crud/data/response"
	request2 "crud/request"
)

type BookService interface {
	Create(ctx context.Context, requst request2.BookCreateRequest)
	Update(ctx context.Context, requst request2.BookUpdateRequest)
	Delete(ctx context.Context, bookId int)
	FindById(ctx context.Context, bookId int) response.BookResponse
	FindAll(ctx context.Context) []response.BookResponse
}
