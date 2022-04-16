package main

import (
	"log"
	"net/http"

	"github.com/chaitsgithub/goprojects/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)

	// http.Handle("/", r)
	log.Println("Starting bookstore at port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
