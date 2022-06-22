package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getByID)
	router.POST("/albums", postAlbum)

	router.Run("localhost:8080")
}

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Lolamsan", Artist: "Jahongir Atajonov", Price: 33},
	{ID: "2", Title: "Gulim", Artist: "Manzura", Price: 32},
	{ID: "3", Title: "Vatanim", Artist: "Ozodbek Nazarbekov", Price: 23},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbum(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{
		"message": "album topilmadi",
	})
}

func updateAlbum(c *gin.Context) {
	// params := c.Param()
}

func deleteAlbum(c *gin.Context) {

}

func searchAlbum(c *gin.Context) {

}
