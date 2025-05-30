package main 


import (
	"net/http"
	"github.com/gin-gonic/gin"
	"errors"
 )

type book struct{
	ID string `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Quantity int `json:"quatity"`
}

var books = []book{
	{ID: "1", Title: "In Search of the Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Greate Gatsby", Author: "F. Scoutt Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
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

func bookById(c *gin.Context){
	id := c.Param("id")
	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return 
	}
	c.IndentedJSON(http.StatusOK, book)
}

func getBookById(id string) (*book, error){

	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")
}

func checkoutBook(c *gin.Context){
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query paramenter."})
		return 
	}

	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return 
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message" : "book not available"})
		return 
	}

	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)
}

func returnBook(c *gin.Context){
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query paramenter."})
		return
	}

	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
	}

	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)
}

func main() {
   router := gin.Default()
   router.GET("/books", getBooks)
   router.GET("books/:id", bookById)
   router.POST("/books", createBook)
   router.PATCH("/checkout", checkoutBook)
   router.PATCH("/return", returnBook)
   router.Run("localhost:8080")
}


