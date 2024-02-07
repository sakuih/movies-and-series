package main

import (
  "net/http"
  "os"
  //"fmt"
  "log"

  "github.com/gin-gonic/contrib/static"
  "github.com/gin-gonic/gin"
  "github.com/joho/godotenv"
)

func main () {
  err := godotenv.Load("local.env")
  if err != nil {
    log.Fatal("Envirioment file not found %s", err)
  }

  valsecret := os.Getenv("SECRET")

  router := gin.Default()

  router.Use(static.Serve("/", static.LocalFile("./views/dist/", true)))

  
  api := router.Group("/api")
  {
    api.GET("/healthCheck", func(c *gin.Context) {
      c.JSON(http.StatusOK, gin.H {
        "message": "pong",
      })
    })


  }


  api.GET("/", getMoviesHandler)
  api.GET("/data", func (c *gin.Context) {
    data := map[string]string{
                "Secret": valsecret,
            }
    c.JSON(200, data)
  })
  api.POST("/jokes/like/:movieID", LikeMovie)
  api.DELETE("/Delete", deleteLikeHandler)

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

