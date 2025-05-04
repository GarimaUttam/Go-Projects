package main 


import (
	"net/http"
	"github.com/gin-gonic/gin"
// 	"errors"
 )

type book struct{
	ID string `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Quatity int `json:"quatity"`
}

var books = []book{
	{ID: "1", Title: "In Search of the Lost Time", Author: "Marcel Proust", Quatity: 2},
	{ID: "2", Title: "The Greate Gatsby", Author: "F. Scoutt Fitzgerald", Quatity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quatity: 6},
}


func getBooks(c *gin.Context){
	c.IndentedJSON(http.StatusOK, books)
}

// Returning does NOT automatically return a response. The BindJSON() method is what handle sending the error response
func createBook(c *gin.Context){
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return 
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func main() {
   router := gin.Default()
   router.GET("/books", getBooks)
   router.POST("/books", createBook)
   router.Run("localhost:8080")
}