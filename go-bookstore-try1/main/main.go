package main

import (
	"chaitsgithub/goprojects/go-bookstore-try1/pkg/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)

	log.Println("Starting bookstore Server...")
	log.Fatal(http.ListenAndServe(":8080",r))
}