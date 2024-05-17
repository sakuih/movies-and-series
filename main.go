package main

import (
  //"net/http"
  //"os"
  //"fmt"
  //"log"

  "movies-and-series.com/src/models"
  "movies-and-series.com/src/routes"
  //"movies-and-series.com/src/controllers"
  //"movies-and-series.com/src/middlewares"
  "movies-and-series.com/src/utils"
  "github.com/gin-gonic/contrib/static"
  //"github.com/gin-gonic/gin"
  //"github.com/joho/godotenv"
)

func main () {
  utils.LoadEnv()
  models.OpenDatabaseConnection()
  //models.AutoMigrateModels()
  router := routes.SetupRoutes()
  //middlewares.RegisterMiddlewares(router)


  router.Use(static.Serve("/", static.LocalFile("./views/dist/", true)))


  router.Run(":3002")
}


