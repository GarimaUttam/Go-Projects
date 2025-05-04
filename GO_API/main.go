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

func main() {
   router := gin.Default()
   router.GET("/books", getBooks)
   router.Run("localhost:8080")
}