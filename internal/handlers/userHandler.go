package handlers

import (
	"net/http"
	"strconv"

	"github.com/Richard-Persson/SAP-Server-API/db"
	"github.com/Richard-Persson/SAP-Server-API/internal/models"
	"github.com/Richard-Persson/SAP-Server-API/internal/tools"
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
	idString := context.Param("id")
	idNumber, parseErr := strconv.ParseInt(idString, 0, 64)

	if parseErr != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Error": parseErr.Error()})
	}

	println("ID: ", idNumber)

	const query = 
		`
		SELECT id, email, first_name, last_name, mobile, billing_code_id
		FROM users 
		WHERE id = $1
		`
	const timeEntriesQuery= `
		SELECT * 
		FROM time_entries 
		WHERE user_id = $1
		`
	var timeEntries []models.TimeEntry 

	dbErr := db.DB.Select(&timeEntries,timeEntriesQuery,idNumber)
	if dbErr != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Error": dbErr.Error()})
		return
	}

	err := db.DB.Get(&user, query,idNumber)
	user.Entries = &timeEntries

	//Removes T00:00:00Z From Date attribute
	tools.RemoveTZ(&timeEntries)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, user)

}
