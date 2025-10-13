package handlers

import (
	"net/http"

	"github.com/Richard-Persson/SAP-Server-API/db"
	"github.com/Richard-Persson/SAP-Server-API/internal/models"
	"github.com/Richard-Persson/SAP-Server-API/internal/payload/requests"
	"github.com/Richard-Persson/SAP-Server-API/internal/tools"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func login(context *gin.Context) {

	var loginRequest requests.LoginRequest

	requestError := context.ShouldBindJSON(&loginRequest);
  if requestError != nil {
    context.JSON(http.StatusBadRequest, gin.H{"error": requestError.Error()})
    return
  }

  var user models.User

  userError := db.DB.Get(&user, "SELECT * FROM users WHERE email=$1", loginRequest.Email)
  if userError != nil {
    context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
    return
  }

	comparison := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)) 

	// If the compared password dosen't match return
  if comparison != nil {
    context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
    return
  }

	const timeEntriesQuery= `
		SELECT * 
		FROM time_entries 
		WHERE user_id = $1
		`
	var timeEntries []models.TimeEntry 

	dbErr := db.DB.Select(&timeEntries,timeEntriesQuery,user.ID)

	user.Entries = &timeEntries

	//Removes T00:00:00Z From Date attribute
	tools.RemoveTZ(&timeEntries)

	if dbErr != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Error": dbErr.Error()})
		return
	}


	//TODO Add Cookie and/or BasicAuthDb() for better authentication and user experience




	context.JSON(http.StatusOK, user)
}
