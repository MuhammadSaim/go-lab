package models

import (
	"time"

	"github.com/MuhammadSaim/go-lab/go-bookstore/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string    `json:"name"`
	Author      string    `json:"author"`
	Publication string    `json:"publication"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.Create(b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookByID(ID int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Find(&getBook, ID)
	return &getBook, db
}

func DeleteBook(ID int64) Book {
	var book Book
	db.Delete(&book, ID)
	return book
}
