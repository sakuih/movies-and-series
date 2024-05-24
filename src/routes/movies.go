package routes

import (
	"github.com/gin-gonic/gin"
	"movies-and-series.com/src/controllers"
	middleware "movies-and-series.com/src/middlewares"
)

func favoritesGroupRouter(router *gin.RouterGroup) {
	api := router.Group("/data/api")
	{
		api.GET("/data", controllers.GetData)
		api.GET("/get", controllers.GetFavorites)
		api.POST("/add", controllers.AddFavorite)
		api.DELETE("/remove", controllers.RemoveFavorite)
	}

}

func SetupRoutes() *gin.Engine {
	r := gin.New()
	r.Use(middleware.Logger(), gin.Recovery(), middleware.CorsMiddleware())
	versionRouter := r.Group("api/v1")
	favoritesGroupRouter(versionRouter)

	return r
}
