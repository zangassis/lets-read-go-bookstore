package models

import (
	"github.com/jinzhu/gorm"
	"github.com/zangassis/lets-read-go-bookstore/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
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

func FindBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func FindBookById(Id int64) (*Book, *gorm.DB) {
	var findBook Book
	db := db.Where("ID=?", Id).Find(findBook)
	return &findBook, db
}

func DeleteBook(Id int64) Book {
	var book Book
	db.Where("ID=?", Id).Delete(book)
	return book
}
