package main 

import (
	"net/http"
  "sync"

  models "app/models"
  utils "app/utils"

  "golang.org/x/crypto/bcrypt"
	"github.com/gin-gonic/gin"
  )

var userStore = map[string]string{}
var revokedTokens = map[string]bool{}
var tokenMutext = &sync.Mutex{}


type UserCredentials struct {
  Name     string `json:"name" binding:"required"`
  Email    string `json:"email" binding:"required,email"`
  Password string `json:"password" binding:"required"`
}

type LoginInput struct {
  Email    string `json:"email" binding:"required,email"`
  Password string `json:"password" binding:"required"`
}

func main() {    
  // Load env.
  utils.LoadEnv()
  //
  host := utils.GetEnv("HOST")
  port := utils.GetEnv("PORT")
  // redis_host := utils.GetEnv("REDIS_HOST")
  // redis_password := utils.GetEnv("REDIS_PASSWORD")
  // jwt_secret := []byte(utils.GetEnv("JWT_KEY"))
  db_dsn := utils.GetEnv("DB_DSN")

  // configuration redis-cache.
  // utils.InitializeConfig(redis_host, redis_password, 0, 2)

  // configuring database connection and migration.
  utils.GetDBConnection(db_dsn)
  utils.Migrate()

  // Routes setup
  r := gin.Default()
  r.POST("/signup", signupHandler)
  r.Run(host + ":" + port)

}


// Signup handler
func signupHandler(c *gin.Context) {
  var creds UserCredentials;
  if err := c.ShouldBindJSON(&creds); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
    return
  }

  passwordHash, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "error in generating password hash"})
  }
  user := models.User{Name: creds.Name, Email: creds.Email, PasswordHash: string(passwordHash)}
	lastID, err := models.CreateUser(utils.DBInstance, user)
	if err != nil {
    c.JSON(http.StatusConflict, gin.H{"error": "User already exsists"})
    return
	}

  c.JSON(http.StatusCreated, gin.H{"message": "User created successfuly", "UserId": lastID})
}


// SIgnIn Handler
func signinHandler(c *gin.Context) {
  var creds LoginInput
  if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
  }

  user, err := models.GetUserByEmail(creds.Email)
  if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User with this email-id does not exsists"})
		return
  }


}
