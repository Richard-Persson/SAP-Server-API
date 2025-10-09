package handlers

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/Richard-Persson/SAP-Server-API/db"
	"github.com/Richard-Persson/SAP-Server-API/internal/models"
)

func listUsers(context *gin.Context) {
  var users []models.User
  if err := db.DB.Select(&users, "SELECT id, email, first_name, last_name , mobile FROM users ORDER BY id DESC"); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }
	context.JSON(http.StatusOK, users)
}

func getUser(context *gin.Context){

	var user models.User
	idString := context.Param("id")
	idNumber, parseErr := strconv.ParseInt(idString, 0, 64)

	if parseErr != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Error": parseErr.Error()})
	}

	println("ID: ", idNumber)

	const query = 
		`
		SELECT id, email, first_name, last_name, mobile 
		FROM users 
		WHERE id = $1
		`

	err := db.DB.Get(&user, query,idNumber)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, user)

}
