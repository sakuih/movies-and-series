package main

import (

  "movies-and-series.com/src/models"
  "movies-and-series.com/src/routes"
  "movies-and-series.com/src/utils"
  "github.com/gin-gonic/contrib/static"
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


