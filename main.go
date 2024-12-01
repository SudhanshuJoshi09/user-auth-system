package main 

import (
	"net/http"
  "sync"
  "fmt"
  "time"

  models "app/models"
  utils "app/utils"

  "golang.org/x/crypto/bcrypt"
	"github.com/gin-gonic/gin"
  )

var userStore = map[string]string{}
var revokedTokens = map[string]bool{}
var tokenMutext = &sync.Mutex{}
var jwtSecret = ""


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
  redis_host := utils.GetEnv("REDIS_HOST")
  redis_password := utils.GetEnv("REDIS_PASSWORD")
  jwtSecret = utils.GetEnv("JWT_KEY")
  db_dsn := utils.GetEnv("DB_DSN")

  // configuration redis-cache.
  utils.InitializeConfig(redis_host, redis_password, 0, 2)

  // configuring database connection and migration.
  utils.GetDBConnection(db_dsn)
  utils.Migrate()

  // Routes setup
  r := gin.Default()
  r.POST("/signup", signupHandler)
  r.POST("/signin", signinHandler)
  r.POST("/signout", signoutHandler)
  r.POST("/refresh", refreshTokenHandler)
  r.Run(host + ":" + port)

}


// Signup handler
func signupHandler(c *gin.Context) {
  var creds UserCredentials
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

  token, err := utils.Get(creds.Email)
  if err == nil {
    c.JSON(http.StatusOK, gin.H{"message": "User has already logged in.", "token": token})
    return
  }

  user, err := models.GetUserByEmail(utils.DBInstance, creds.Email)
  if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User with this email-id does not exsists"})
		return
  }

  err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(creds.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

  token, err = utils.GenerateLoginToken(creds.Email, jwtSecret)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "error in generating login token"})
    fmt.Println("aksdjflajksdfk ", err, jwtSecret)
    return
  }

  // Hardcoding this to 10 mins, let's have a config for this. improve the config mgmt, make it globally available
  utils.Set(user.Email, token, 10*time.Minute)

  c.JSON(http.StatusOK, gin.H{"message": "User logged in successfuly", "token": token})
  return
}


// SignOut Handler
func signoutHandler(c *gin.Context) {

  authHeader := c.GetHeader("Authorization")
  email, err := utils.DecodeLoginToken(authHeader, jwtSecret)
  if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request, unable to decode your token"})
    return
  }

  _, err = utils.Get(email)
  if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Your Session Token has expired or the user does not exists, Please login back again"})
    return
  }
  err = utils.Delete(email)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "error in deleting login token"})
    return
  }
  c.JSON(http.StatusOK, gin.H{"message": "User has logged out successfuly"})
  return
}

func refreshTokenHandler(c *gin.Context) {
  authHeader := c.GetHeader("Authorization")
  email, err := utils.DecodeLoginToken(authHeader, jwtSecret)
  if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request, unable to decode your token"})
    return
  }

  _, err = utils.Get(email)
  if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Your Session Token has expired or the user does not exists, Please login back again"})
    return
  }
  token, err := utils.GenerateLoginToken(email, jwtSecret)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "error in renewing login token"})
    return
  }

  utils.Set(email, token, 10*time.Minute)
  c.JSON(http.StatusOK, gin.H{"message": "Your token has been refreshed."})
  return
}
