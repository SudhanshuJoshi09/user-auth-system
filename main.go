package main 

import (
  // "net/http"
  // // "sync"
  // // "time"
  // "fmt"
  cutils "app/utils"
  )

func main() {    
  cutils.LoadEnv()

  host := cutils.GetEnv("HOST")
  port := cutils.GetEnv("PORT")
  redis_host := cutils.GetEnv("REDIS_HOST")
  redis_password := cutils.GetEnv("REDIS_PASSWORD")
  jwt_secret := []byte(cutils.GetEnv("JWT_KEY"))
  db_dsn := cutils.GetEnv("DB_DSN")

  cutils.InitializeConfig(redis_host, redis_password, 0, 2)

// }

// import (
//   "database/sql"
//   _ "github.com/go-sql-driver/mysql"
//   // "gorm.io/driver/mysql"
//   // "gorm.io/gorm"
//   "fmt"
// )
//
// func main() {
//   // sqlDB, err := sql.Open("mysql", "")
//   // fmt.Println("akfjsdlkfjasd ", err)
//   // gormDB, err := gorm.Open(mysql.New(mysql.Config{
//   //   Conn: sqlDB,
//   // }), &gorm.Config{})
//   // fmt.Println("akfjsdlkfjasd ", err)
//   // fmt.Println("akfjsdlkfjasd ", gormDB)
//
//   sqlDB, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3307)/udb")
//   err = sqlDB.Ping()
//   if err != nil {
//     fmt.Println("dklfajdlsfjka ", err)
//   }
//   defer sqlDB.Close()
// }


// import (
//   "fmt"
//   models "app/models"
//   utils "app/utils"
// )
//
// func main() {
// 	// Initialize database connection using Singleton pattern
// 	dsn := "root:@tcp(127.0.0.1:3307)/udb"
// 	db := utils.GetDBConnection(dsn)
//
// 	// Run migrations
// 	utils.Migrate()
//
// 	// Create a new user
// 	user := models.User{Name: "Sudhanshu", Email: "sudhanshujoshi49@gmail.com"}
// 	lastID, err := models.CreateUser(db, user)
// 	if err != nil {
// 		fmt.Println("Error creating user:", err)
// 	} else {
// 		fmt.Println("User created with ID:", lastID)
// 	}
//
// 	// Get the user by ID
// 	// fetchedUser, err := models.GetUserByID(db, lastID)
// 	//
// 	// if err != nil {
// 	// 	fmt.Println("Error fetching user:", err)
// 	// } else {
// 	// 	fmt.Println("Fetched user:", fetchedUser)
// 	// }
// 	//
// 	// // Update the user
// 	// fetchedUser.Name = "John Updated"
// 	// err = models.UpdateUser(db, fetchedUser)
// 	// if err != nil {
// 	// 	fmt.Println("Error updating user:", err)
// 	// } else {
// 	// 	fmt.Println("User updated successfully")
// 	// }
// 	//
// 	// // Delete the user
// 	// err = models.DeleteUser(db, lastID)
// 	// if err != nil {
// 	// 	fmt.Println("Error deleting user:", err)
// 	// } else {
// 	// 	fmt.Println("User deleted successfully")
// 	// }
// }

