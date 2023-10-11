package Service

import (
	"context"
	"crud/data/response"
	"crud/helper"
	"crud/model"
	"crud/repository"
	request2 "crud/request"
)

type BookServiceImpl struct {
	BookRepository repository.BookRespository
}

func NewBookServicrImpl(bookRepository repository.BookRespository) BookService {
	return &BookServiceImpl{BookRepository: bookRepository}
}

func (b *BookServiceImpl) Create(ctx context.Context, request request2.BookCreateRequest) {

	book := model.Book{
		Name: request.Name,
	}

	b.BookRepository.Save(ctx, book)
}

func (b *BookServiceImpl) Update(ctx context.Context, request request2.BookUpdateRequest) {

	book, err := b.BookRepository.FindById(ctx, request.Id)
	helper.PanicIfErr(err)
	book.Name = request.Name
	b.BookRepository.Update(ctx, book)

}

func (b *BookServiceImpl) Delete(ctx context.Context, bookId int) {
	book, err := b.BookRepository.FindById(ctx, bookId)
	helper.PanicIfErr(err)
	b.BookRepository.Delete(ctx, book.Id)

}

func (b *BookServiceImpl) FindById(ctx context.Context, bookId int) response.BookResponse {

	book, err := b.BookRepository.FindById(ctx, bookId)
	helper.PanicIfErr(err)
	return response.BookResponse(book)

}

func (b *BookServiceImpl) FindAll(ctx context.Context) []response.BookResponse {
	books := b.BookRepository.FindAll(ctx)
	var bookResp []response.BookResponse

	for _, value := range books {
		book := response.BookResponse{Id: value.Id, Name: value.Name}
		bookResp = append(bookResp, book)
	}

	return bookResp
}
