package handlers

import (
  "net/http"
	"time"

	"github.com/gin-contrib/cors"
  "github.com/gin-gonic/gin"
  "github.com/Richard-Persson/SAP-Server-API/db"
  "github.com/Richard-Persson/SAP-Server-API/models"
)

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
  return router
}



func listUsers(context *gin.Context) {
  var users []models.User
  if err := db.DB.Select(&users, "SELECT id, full_name FROM users ORDER BY id DESC"); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }
	context.JSON(http.StatusOK, users)
}
