package models

import (
  "gorm.io/gorm"
)

type Favorite struct {
  gorm.Model
  movieID       int
  movieTitle    string
  User          User
}

type User struct {
  gorm.Model
  userID        int
  username      string
  password      string
}

func (favorite *Favorite) Save() (*Favorite, error) {
  err := Database.Model(&favorite).Create(&favorite).Error
  if err != nil {
    return &Favorite{}, err
  }
  return favorite, nil
}

func FetchAllFavorite() (*[]Favorite, error) {
  var favorites []Favorite
  err := Database.Find(&favorites).Error
  if err != nil {
    return &[]Favorite{}, err
  }
  return &favorites, nil
}

func DeleteStartup(id string= error {
 err := Database.Model(&Favorite{}).Where("id = ?", id).Delete(&Favorite{}).Error
  if err != nil {
    return err
   }
  return nil
}



