package models

import (
	"github.com/jinzhu/gorm"
	"github.com/shivamgutgutia/go-bookstore/pkg/config"
	_"github.com/jinzhu/gorm/dialects/mysql"
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

func (book *Book) CreateBook() *Book {
	db.NewRecord(book)
	db.Create(&book)
	return book
}

func GetAllBooks() []Book{
	var books []Book
	db.Find(&books)
	return books
}

func GetBookById(Id int64)(*Book, *gorm.DB){
	var getBook Book
	db := db.Where("ID=?",Id).Find(&getBook)
	return &getBook,db
}

func DeleteBook(Id int64) Book {
	var book Book
	db.Where("ID=?",Id).Delete(book)
	return book
}


