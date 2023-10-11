package controller

import (
	"crud/Service"
	"crud/data/response"
	"crud/helper"
	"crud/request"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type BookController struct {
	BookService Service.BookService
}

func NewBookController(bookService Service.BookService) *BookController {
	return &BookController{BookService: bookService}
}

func (Controller *BookController) Create(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	bookCreateRequest := request.BookCreateRequest{}
	helper.ReadRequestBody(requests, &bookCreateRequest)
	Controller.BookService.Create(requests.Context(), bookCreateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (Controller *BookController) Update(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {

	bookUpdateRequest := request.BookUpdateRequest{}
	helper.ReadRequestBody(requests, &bookUpdateRequest)
	bookId := params.ByName("bookId")
	id, err := strconv.Atoi(bookId)
	helper.PanicIfErr(err)
	bookUpdateRequest.Id = id
	Controller.BookService.Update(requests.Context(), bookUpdateRequest)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}
	helper.WriteResponseBody(writer, webResponse)
}

func (Controller *BookController) Delete(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {

	bookId := params.ByName("bookId")
	id, err := strconv.Atoi(bookId)
	helper.PanicIfErr(err)
	Controller.BookService.Delete(requests.Context(), id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}
	helper.WriteResponseBody(writer, webResponse)
}

func (Controller *BookController) FindAll(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {

	result := Controller.BookService.FindAll(requests.Context())
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}
	helper.WriteResponseBody(writer, webResponse)
}

func (Controller *BookController) FindById(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {

	bookId := params.ByName("BookId")
	id, err := strconv.Atoi(bookId)
	helper.PanicIfErr(err)
	result := Controller.BookService.FindById(requests.Context(), id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helper.WriteResponseBody(writer, webResponse)
}
