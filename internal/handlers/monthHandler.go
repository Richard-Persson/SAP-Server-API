package handlers

import (
	"net/http"

	"github.com/Richard-Persson/SAP-Server-API/db"
	"github.com/Richard-Persson/SAP-Server-API/internal/models"
	"github.com/gin-gonic/gin"
)



func listMonths(context *gin.Context){

	var months[] models.Month

	const query = 
		`
		SELECT id, year, month, user_id, total_hours FROM months 

		`

	err := db.DB.Select(&months, query)

	if err == nil {
	context.JSON(http.StatusInternalServerError, gin.H{"Error" : err.Error()})
		return
	}


	context.JSON(http.StatusOK, months)
}
