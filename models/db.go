package models

import (
  "fmt"
  "os"
  "gorm.io/gorm"
  "gorm.io/driver/postgres"
)

var Database *gorm.DB

func openDatabaseConn () {
  var err error
  host := os.Getenv("")
  username := os.Getenv("DB_USER")
  password := os.Getenv("DB_PASS")
  databaseName := os.Getenv("DB_NAME")
  port := os.Getenv("DB_PORT")

  dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require TimeZone=Europe/Helsinki", host, username, password, databaseName, port)
  Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

 if err != nil {
  panic(err)
 } else {
  fmt.Println("lets go")
 }

}

func AutoMigrateModels() {
  Database.AutoMigrate(&Favorite{})
}


