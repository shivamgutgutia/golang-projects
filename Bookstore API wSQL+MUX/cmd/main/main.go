package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shivamgutgutia/go-bookstore/pkg/routers"
)

func main() {
	r := mux.NewRouter()
	routers.BookStoreRouter(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
