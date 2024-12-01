package models

import (
  "gorm.io/gorm"
)

type User struct {
	ID           uint   `gorm:"primaryKey"`
	Name         string `gorm:"not null"`
	Email        string `gorm:"unique;not null"`
  PasswordHash string `gorm:"not null"`
}

func CreateUser(db *gorm.DB, user User) (uint, error) {
	result := db.Create(&user)
	return user.ID, result.Error
}

func GetUserByID(db *gorm.DB, id uint) (User, error) {
	var user User
	result := db.First(&user, id)
	return user, result.Error
}

func GetUserByEmail(db *gorm.DB, email string) (User, error) {
	var user User
	result := db.First(&user, "email = ?", email)
	return user, result.Error
}


func UpdateUser(db *gorm.DB, user User) error {
	result := db.Save(&user)
	return result.Error
}

func DeleteUser(db *gorm.DB, id uint) error {
	result := db.Delete(&User{}, id)
	return result.Error
}

