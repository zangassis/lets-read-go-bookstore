package routes

import (
	"github.com/gorilla/mux"
	"github.com/zangassis/lets-read-go-bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book", controllers.FindBooks).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.FindBookById).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/ook/{ID}", controllers.DeleteBook).Methods("DELETE")
}
