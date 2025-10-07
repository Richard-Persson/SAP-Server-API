package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Richard-Persson/SAP-Server-API/db"
	"github.com/Richard-Persson/SAP-Server-API/internal/models"
	"github.com/Richard-Persson/SAP-Server-API/internal/payload/requests"
	"github.com/Richard-Persson/SAP-Server-API/internal/tools"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// buildAccounts queries all users and returns a gin.Accounts map.

func buildAccounts() gin.Accounts {
	var users []models.User // struct { email, Password string }
	if err := db.DB.Select(&users,
		"SELECT email, password FROM users"); err != nil {
		log.Fatalf("failed to load users for basic auth: %v", err)
	}

	accounts := gin.Accounts{}
	for _, u := range users {
		accounts[u.Email] = u.Password
	}
	return accounts
}

func Router() http.Handler {
	router := gin.Default()


	  // CORS config for React dev server
  router.Use(cors.New(cors.Config{
    AllowOrigins:     []string{"http://localhost:5173"},
    AllowMethods:     []string{"GET", "POST", "OPTIONS"},
    AllowHeaders:     []string{"Content-Type"},
    AllowCredentials: true,
    MaxAge:           12 * time.Hour,
  }))


	router.GET("/users",listUsers)
	router.POST("/login", login)
	router.POST("/register", register)


	authGroup := router.Group("/")
	authGroup.Use(gin.BasicAuth(buildAccounts()))
  return router
}


func register(c *gin.Context){

  var request requests.RegisterRequest 
  if err := c.ShouldBindJSON(&request); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  // Hash the plaintext password
  hash, err := tools.HashPassword(request.Password)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
    return
  }

  // Insert user into the database
  const query = `
		INSERT INTO users (email, first_name, last_name, mobile, password)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING id, first_name, last_name, email, mobile, password
  `
  var user models.User
  err = db.DB.Get(&user, query, request.Email, request.FirstName, request.LastName, request.Mobile,hash)

  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create user"})
		fmt.Printf(" %v", err)
    return
  }

	//Returns the user which was created
	c.JSON(http.StatusCreated, user)
}


func login(context *gin.Context) {

	var request requests.LoginRequest

	requestError := context.ShouldBindJSON(&request);
  if requestError != nil {
    context.JSON(http.StatusBadRequest, gin.H{"error": requestError.Error()})
    return
  }

  var user models.User

  userError := db.DB.Get(&user, "SELECT id, email, password FROM users WHERE email=$1", request.Email)
  if userError != nil {
    context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
    return
  }

	comparison := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) 

	// If the compared password dosen't match return
  if comparison != nil {
    context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
    return
  }

	//TODO Add Cookie and/or BasicAuthDb() for better authentication and user experience

	context.Redirect(http.StatusFound, "/home")

}

func listUsers(context *gin.Context) {
  var users []models.User
  if err := db.DB.Select(&users, "SELECT id, first_name, last_name FROM users ORDER BY id DESC"); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }
	context.JSON(http.StatusOK, users)
}
