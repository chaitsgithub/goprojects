package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request#1")
		w.Write([]byte("Hello World!"))
	})

	http.HandleFunc("/welcome", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request#2")
		w.Write([]byte("Welcome to Go Server!"))
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request#3")
		w.Write([]byte("Default Page!"))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
