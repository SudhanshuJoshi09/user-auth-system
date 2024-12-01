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
	DBInstance *gorm.DB
	dbOnce     sync.Once
)

func GetDBConnection(dsn string) {
	dbOnce.Do(func() {
		var err error
		DBInstance, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
      log.Fatal("Error opening DB connection:", err)
			panic(err)
		}
		fmt.Println("Database connection established")
	})
}

func Migrate() {
	err := DBInstance.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Error during migration:", err)
		panic(err)
	}
	fmt.Println("Migrations applied successfully")
}

