package main

import (
  "net/http"
  "os"
  //"fmt"
  "log"

  "movies-and-series.com/src/models"
  "movies-and-series.com/src/routes"
  //"movies-and-series.com/src/controllers"
  "movies-and-series.com/src/middlewares"

  "github.com/gin-gonic/contrib/static"
  "github.com/gin-gonic/gin"
  "github.com/joho/godotenv"
)

func main () {
  err := godotenv.Load("local.env")
  if err != nil {
    log.Fatal("Envirioment file not found %s", err)
  }
  models.OpenDatabaseConnection()
  models.AutoMigrateModels()
  router := routes.SetupRoutes()
  middlewares.RegisterMiddlewares(router)

  valsecret := os.Getenv("SECRET")

  //router := gin.Default()

  router.Use(static.Serve("/", static.LocalFile("./views/dist/", true)))

  api := router.Group("/api")
  {
    api.GET("/data", func (c *gin.Context) {
      data := map[string]string{
                  "Secret": valsecret,
              }
      c.JSON(200, data)
    })

    api.GET("/", getMoviesHandler)
    api.POST("/like/:movieID", LikeMovie)
    api.DELETE("/Delete", deleteLikeHandler)

  }



  router.Run(":3002")
}

// Routes

func deleteLikeHandler (c *gin.Context) {
  c.Header("Content-Type", "application/json")

  c.JSON(http.StatusOK, gin.H{
    "message": "delete",
  })
}

func getMoviesHandler (c *gin.Context) {
  c.Header("Content-Type", "application/json")

  c.JSON(http.StatusOK, gin.H {
    "message":"Jokes handler not implemented yet",
  })
}

//func getApiSecret (c *gin.Context, valsecret) {}

func LikeMovie(c *gin.Context) {
  c.Header("Content-Type", "application/json")
  c.JSON(http.StatusOK, gin.H {
    "message": "LikeJoke handler not implemented yet",
  })
}

