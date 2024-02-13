package routes


import (
  "github.com/gin-gonic/gin"
  //"movies-and-series.com/src/routes"
  "movies-and-series.com/src/controllers"
)

func favoritesGroupRouter (router *gin.RouterGroup) {
  favor := router.Group("/favor")

  favor.GET("/", controllers.GetFavorites)
  favor.POST("/add", controllers.AddFavorite)
  favor.DELETE("/remove", controllers.RemoveFavorite)

}

func SetupRoutes () *gin.Engine {
  r := gin.Default()
  versionRouter := r.Group("api/v1")
  favoritesGroupRouter(versionRouter)

  return r
}

