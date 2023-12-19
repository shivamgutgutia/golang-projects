package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/shivamgutgutia/go-bookstore/pkg/models"
	"github.com/shivamgutgutia/go-bookstore/pkg/utils"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		panic(err)
	}
	bookDetails,_ :=models.GetBookById(Id)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	createBook := &models.Book{}
	utils.ParseBody(r, createBook)
	book :=  createBook.CreateBook()
	res,_ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id,err := strconv.ParseInt(bookId,0,0)
	if err !=nil{
		panic(err)
	}
	book := models.DeleteBook(Id)
	res,_ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request){
	updateBook := models.Book{}
	utils.ParseBody(r,updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id,err := strconv.ParseInt(bookId,0,0)
	if err !=nil{
		panic(err)
	}
	book,db := models.GetBookById(Id)
	if updateBook.Name!=""{
		book.Name=updateBook.Name
	}
	if updateBook.Author!=""{
		book.Author=updateBook.Author
	}
	if updateBook.Publication!=""{
		book.Publication=updateBook.Publication
	}
	db.Save(&book)
	res,_ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
