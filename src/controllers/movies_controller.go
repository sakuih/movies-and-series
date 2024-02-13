  package controllers

import (

  "net/http"
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
func RemoveFavorite(c *gin.Context){
}

