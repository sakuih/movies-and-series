package routes


import (
  "github.com/gin-gonic/gin"
  //"movies-and-series.com/src/routes"
  "movies-and-series.com/src/controllers"
  middleware "movies-and-series.com/src/middlewares"
  "os"
)

func favoritesGroupRouter (router *gin.RouterGroup) {
  favor := router.Group("/favor")
  api := router.Group("/api")
  {
    api.GET("/data", func (c *gin.Context) {
      valsecret := os.Getenv("SECRET")
      data := map[string]string{
                  "Secret": valsecret,
              }
      c.JSON(200, data)
    })
  }

  favor.GET("/get", controllers.GetFavorites)
  favor.POST("/add", controllers.AddFavorite)
  favor.DELETE("/remove", controllers.RemoveFavorite)

}

func SetupRoutes () *gin.Engine {
  r := gin.New()
  r.Use(middleware.Logger())
  versionRouter := r.Group("api/v1")
  favoritesGroupRouter(versionRouter)

  return r
}

