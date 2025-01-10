package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/MuhammadSaim/go-lab/go-bookstore/pkg/models"
	"github.com/MuhammadSaim/go-lab/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

// return all the books
func GetBook(w http.ResponseWriter, r *http.Request) {
	// get all the books from the DB
	books := models.GetAllBooks()

	// Marshal the data to json and response it back
	res, _ := json.Marshal(books)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// find one book bt id
func GetBookById(w http.ResponseWriter, r *http.Request) {
	// get all the params from request
	params := mux.Vars(r)

	// get bookId in params
	bookId := params["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	// check there is any error
	if err != nil {
		fmt.Println("Error while parsing")
	}

	// get the book details from DB
	bookDetails, _ := models.GetBookByID(ID)

	// return the response
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// an endpoint to create a book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	// create an empty struct
	CreateBook := &models.Book{}

	// parse the requets body and store in CreateBook
	utils.ParseBody(r, CreateBook)

	// store the book in DB
	b := CreateBook.CreateBook()

	// return the response
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

// Delete the book by the ID
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	// get all the params from request
	params := mux.Vars(r)

	// get bookId in params
	bookId := params["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	// check there is any error
	if err != nil {
		fmt.Println("Error while parsing")
	}

	// delete the book in DB
	book := models.DeleteBook(ID)

	// return the response
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// update the book with the id
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	// set the empty struct
	updateBook := &models.Book{}

	// parse the body
	utils.ParseBody(r, updateBook)

	// get all the params from request
	params := mux.Vars(r)

	// get bookId in params
	bookId := params["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	// check there is any error
	if err != nil {
		fmt.Println("Error while parsing")
	}

	// get book by the id
	bookDetails, db := models.GetBookByID(ID)

	// verify there is no empty values
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}

	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}

	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}

	// save the db
	db.Save(&bookDetails)

	// return the response
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
