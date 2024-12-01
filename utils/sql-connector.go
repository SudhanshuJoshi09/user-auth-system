package utils

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
  "log"
  models "app/models"
)

var (
	dbInstance *gorm.DB
	dbOnce       sync.Once
)

func GetDBConnection(dsn string) *gorm.DB {
	dbOnce.Do(func() {
		var err error
		dbInstance, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
      log.Fatal("Error opening DB connection:", err)
			panic(err)
		}
		fmt.Println("Database connection established")
	})
	return dbInstance
}

func Migrate() {
	err := dbInstance.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Error during migration:", err)
		panic(err)
	}
	fmt.Println("Migrations applied successfully")
}

