package models

import (
	"github.com/chaitsgithub/goprojects/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func (b *Book) GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func (b *Book) GetBookById(id int) (*Book, *gorm.DB) {
	var getBook Book
	d := db.First(&getBook, id)
	return &getBook, d
}

func (b *Book) DeleteBook(id int) *gorm.DB {
	var book Book
	d := db.Where("ID=?", id).Delete(book)
	return d
}
