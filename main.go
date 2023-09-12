package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "The Enigmatic Galaxy", Author: "Lillian Weaver", Quantity: 5},
	{ID: "2", Title: "Chronicles of the Lost Kingdom", Author: "Samuel Northwood", Quantity: 7},
	{ID: "3", Title: "Whispers in the Shadows", Author: "Evelyn Blackwell", Quantity: 3},
	{ID: "4", Title: "The Alchemist's Legacy", Author: "Oliver Stone", Quantity: 6},
	{ID: "5", Title: "Secrets of the Silent Forest", Author: "Sophia Mitchell", Quantity: 4},
	{ID: "6", Title: "The Mysterious Codex", Author: "Alexander Gray", Quantity: 2},
	{ID: "7", Title: "Lost in the Labyrinth", Author: "Victoria Knight", Quantity: 8},
	{ID: "8", Title: "Echoes of Eternity", Author: "Lucas Taylor", Quantity: 10},
	{ID: "9", Title: "The Secret Society", Author: "Isabel Carter", Quantity: 1},
	{ID: "10", Title: "The Quantum Paradox", Author: "Nathaniel Reed", Quantity: 9},
}

func createBook(context *gin.Context) {
	var newBook book
	if err := context.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	context.IndentedJSON(http.StatusCreated, newBook)
}

func getBooks(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, books)
}

func getBookByID(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")
}

func bookById(context *gin.Context) {
	id := context.Param("id")
	book, err := getBookByID(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, book)
}

func checkoutBook(context *gin.Context) {
	id := context.Param("id")
	_, err := getBookByID(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	for i, b := range books {
		if b.ID == id {
			if b.Quantity > 0 {
				books[i].Quantity--
				context.IndentedJSON(http.StatusOK, gin.H{"message": "book checked out"})
			} else {
				context.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not available"})
			}
		}
	}
}

func main() {
	router := gin.Default()

	router.GET("/books", getBooks)
	router.POST("/books/new", createBook)
	router.GET("/books/:id", bookById)
	router.POST("/books/checkout/:id", checkoutBook)
	router.Run("localhost:8080")

}
