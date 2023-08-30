package routes

import (
	"chaitsgithub/goprojects/go-bookstore-try1/pkg/handlers"

	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(r *mux.Router) {
	r.HandleFunc("/", handlers.DefaultHandler).Methods("GET")
	r.HandleFunc("/book", handlers.CreateBook).Methods("POST")
}
