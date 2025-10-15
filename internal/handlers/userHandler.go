package handlers

import (
	"net/http"
	"strconv"

	"github.com/Richard-Persson/SAP-Server-API/db"
	"github.com/Richard-Persson/SAP-Server-API/internal/models"
	"github.com/Richard-Persson/SAP-Server-API/internal/queries"
	"github.com/gin-gonic/gin"
)

func listUsers(context *gin.Context) {
  var users []models.User

  if err := db.DB.Select(&users, "SELECT id, email, first_name, last_name , mobile, billing_code_id FROM users ORDER BY id DESC"); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }
	context.JSON(http.StatusOK, users)
}

func getUserById(context *gin.Context){

	var user models.User
	var timeEntries []models.TimeEntry 

	userId, parseErr := strconv.ParseInt(context.Param("id"), 0, 64)
	if parseErr != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Error": parseErr.Error()})
	}

	const query = 
		`
		SELECT id, email, first_name, last_name, mobile, billing_code_id
		FROM users 
		WHERE id = $1
		`
	err := db.DB.Get(&user, query,userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"User not found: ": err.Error()})
		return
	}

	//Get all entries for a single user
	queries.GetTimeEntriesByUserId(&timeEntries, userId, context)
	user.Entries = &timeEntries

	context.JSON(http.StatusOK, user)

}
