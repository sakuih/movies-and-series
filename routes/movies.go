package routes


import (
  "github.com/gin-gonic/gin"
  //"movies-and-series/routes"
)

func favoritesGroupRouter (router *gin.RouterGroup) {
  favor := router.Group("/favor")

  favor.GET("/", controllers.getFavorites)
  favor.POST("/add", controllers.addFavorite)
  favor.DELETE("/remove", controllers.removeFavorite)

}

func setupRoutes () {
  r := gin.Default()
  versionRouter := r.Group("api/v1")
  favoritesGroupRouter(versionRouter)

  return r
}

