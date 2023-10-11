package main

import (
	"crud/Service"
	"crud/config"
	"crud/controller"
	"crud/helper"
	"crud/repository"
	"crud/router"
	"fmt"
	"net/http"
)

func main() {
	fmt.Printf("Start main")

	//database
	db := config.DatabaseConnection()

	//repository

	bookRepository := repository.NewBookRespository(db)

	//service
	bookService := Service.NewBookServicrImpl(bookRepository)

	//Controller

	bookCotroller := controller.NewBookController(bookService)

	//router
	router := router.NewRouter(bookCotroller)

	server := http.Server{Addr: "localhost:8888", Handler: router}

	err := server.ListenAndServe()

	helper.PanicIfErr(err)

}
