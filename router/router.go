package router

import (
	"crud/controller"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func NewRouter(bookController *controller.BookController) *httprouter.Router {

	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "welcone Home")
	})

	router.GET("/api/book", bookController.FindAll)

	router.GET("/api/book/:bookId", bookController.FindById)

	router.PATCH("/api/book/:bookId", bookController.Update)

	router.POST("/api/book", bookController.Create)

	router.DELETE("/api/book/:bookId", bookController.Delete)

	return router

}
