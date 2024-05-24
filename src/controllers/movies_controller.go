package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"movies-and-series.com/src/models"
)

func GetFavorites(c *gin.Context) {
	favorites, err := models.FetchAllFavorites()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Fetch Succesful", "data": favorites})
}

func AddFavorite(c *gin.Context) {
	var input models.Favorite

	if err := c.ShouldBindJSON(&input); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error(), "data": nil})

		return
	}
	savedFavorite, err := input.Save()

	if err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"status": "failed", "message": err.Error(), "data": nil})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "success", "message": "favorite saved successfully", "data": savedFavorite})

}
func RemoveFavorite(c *gin.Context) {
}

func GetData(c *gin.Context) {

	valsecret := os.Getenv("SECRET")
	searchValue := c.Query("searchValue")
	if searchValue == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name query parameter is required"})
		return
	}

	url := fmt.Sprintf("http://www.omdbapi.com/?s=%s&apikey=%s", valsecret, searchValue)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, resp.Body)

}
