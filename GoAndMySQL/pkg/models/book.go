package models

import (
	// "github.com/jinzhu/gorm"
	 "gorm.io/gorm"
	"github.com/GarimaUttam/Go-Projects/GoAndMySQL/pkg/config"
)

var db *gorm.DB

type Book struct{
	gorm.Model
	Name string `json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

func init(){
	config.Connect()
	db = config.GetDB()
	if db == nil {
		panic("Failed to connect to database")
	}
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book{
	db.NewRecord(b)
	db.Create(&b)
	return b
}