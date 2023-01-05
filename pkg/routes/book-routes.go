package routes

import (
	"github.com/gorilla/mux"
	"github.com/ncalamsyah/go-book-ms/pkg/controllers"
)

var RegisterBookRoutes = func(router *mux.Router) {
	router.HandleFunc("/books", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/books", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/books/{book_id}", controllers.GetBookByID).Methods("GET")
	router.HandleFunc("/books/{book_id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{book_id}", controllers.DeleteBook).Methods("DELETE")
}
