package main

import (
	"net/http"

	"fmt"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Tune jo na kaha", Artist: "Sonu Nigam", Price: 100.00},
	{ID: "2", Title: "Aap hamari jaan ban gye", Artist: "Mithoon", Price: 100.00},
	{ID: "3", Title: "Tere bina jina saja ho gya", Artist: "Shreya Ghosal", Price: 100.00},
}

// REST API's
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbumById(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("ID -", id)

	for _, album := range albums {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func main() {
	fmt.Println("---------- Welcome to first rest api ----------")

	fmt.Println(albums)
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8000")
}
