package controllers

import (

  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/username/go-gin-postgresql-backend/src/models"
)

func getFavorites(c *gin.Context) {
  favorites, err := models.FetchAllFavorites()
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
    return
  }
  c.JSON(http.StatusOK, gin.H{"message": "Fetch Succesful"})
}

func addFavorite(c *gin.Context) {
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
func removeFavorite(c *gin.Context){
}

