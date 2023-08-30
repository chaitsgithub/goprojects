package handlers

import (
	"chaitsgithub/goprojects/go-bookstore-try1/pkg/models"
	"encoding/json"
	"net/http"
)

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	resp := "Triggered Default End Point"
	w.Write([]byte(resp))
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var NewBook models.Book
	req := json.Unmarshal([]byte(r.Body), &NewBook)

	w.Write([]byte(resp))
}
