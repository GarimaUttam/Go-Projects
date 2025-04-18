package models

import (
	"gorm.io/gorm"
	"github.com/GarimaUttam/Go-Projects/GoAndMySQL/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	if db == nil {
		panic("Failed to connect to database")
	}
	db.AutoMigrate(&Book{}) // This is fine, but avoid calling it repeatedly in production
}

func (b *Book) CreateBook() *Book {
	result := db.Create(&b)
	if result.Error != nil {
		panic("Failed to create book:" + result.Error.Error())
	}
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	result := db.Where("id = ?", Id).Find(&getBook)
	return &getBook, result
}

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID = ?", ID).Delete(&book)
	return book
}
