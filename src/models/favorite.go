package models

import (
  "gorm.io/gorm"
  "time"
)

type Favorite struct {
  gorm.Model
  movieID       int
  movieTitle    string
  User          User
  createdAt     time.Time
  UpdatedAt     time.Time
}

func (favorite *Favorite) Save() (*Favorite, error) {
  err := Database.Model(&favorite).Create(&favorite).Error
  if err != nil {
    return &Favorite{}, err
  }
  return favorite, nil
}

func FetchAllFavorites() (*[]Favorite, error) {
  var favorites []Favorite
  err := Database.Find(&favorites).Error
  if err != nil {
    return &[]Favorite{}, err
  }
  return &favorites, nil
}

func DeleteStartup(id string) error {
 err := Database.Model(&Favorite{}).Where("id = ?", id).Delete(&Favorite{}).Error
  if err != nil {
    return err
   }
  return nil
}



