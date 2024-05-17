
package utils

import (
  "log"

  "github.com/joho/godotenv"
)

func LoadEnv() {
  err := godotenv.Load("./local.env")
  if err != nil {
    log.Fatal(".env file not found")
  }
}
