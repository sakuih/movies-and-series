package models

import (
  "fmt"
  "os"
  "gorm.io/gorm"
  "gorm.io/driver/postgres"
)

var Database *gorm.DB

func OpenDatabaseConnection () {
  var err error
  host := os.Getenv("DB_HOST")
  username := os.Getenv("DB_USER")
  password := os.Getenv("DB_PASS")
  databaseName := os.Getenv("DB_NAME")
  port := os.Getenv("DB_PORT")

  dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, username, password, databaseName, port)
  //dsn := fmt.Sprintf("host=%s", host)
  Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

 if err != nil {
  panic(err)
 } else {
  fmt.Println("Connected to database")
 }

}

/*func AutoMigrateModels() {
  Database.AutoMigrate(&Favorite{})
}
*/


