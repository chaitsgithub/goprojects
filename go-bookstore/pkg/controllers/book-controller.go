package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/chaitsgithub/goprojects/go-bookstore/pkg/models"
	"github.com/chaitsgithub/goprojects/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	book := CreateBook.CreateBook()

	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := NewBook.GetAllBooks()
	res, _ := json.Marshal(newBooks)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["bookId"]

	w.Header().Set("Content-Type", "application/json")
	ID, err := strconv.Atoi(bookId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	bookDetails, db := NewBook.GetBookById(ID)
	if db.RecordNotFound() {
		http.Error(w, "bookID not found", http.StatusNotFound)
		return
	}

	res, _ := json.Marshal(bookDetails)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["bookId"]

	ID, err := strconv.Atoi(bookId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	bookDetails, db := NewBook.GetBookById(ID)
	var updateBook models.Book
	utils.ParseBody(r, updateBook)

	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)

	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["bookId"]

	ID, err := strconv.Atoi(bookId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	NewBook.DeleteBook(ID)
	newBooks := NewBook.GetAllBooks()

	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
