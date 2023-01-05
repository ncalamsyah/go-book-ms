package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/ncalamsyah/go-book-ms/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookRoutes(r)

	fmt.Println("server running on port 8000")
	log.Fatal(http.ListenAndServe("localhost:8000", r))
}
