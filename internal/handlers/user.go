package handlers

import (
  "net/http"

  "github.com/gin-gonic/gin"
  "github.com/Richard-Persson/SAP-Server-API/db"
  "github.com/Richard-Persson/SAP-Server-API/models"
)

func Router() http.Handler {
	router := gin.Default()
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
