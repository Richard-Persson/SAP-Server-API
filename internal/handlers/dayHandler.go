package handlers

import (
	"net/http"

	"github.com/Richard-Persson/SAP-Server-API/db"
	"github.com/Richard-Persson/SAP-Server-API/internal/models"
	"github.com/gin-gonic/gin"
)





func getDaysByUserId(context *gin.Context){

	userId := context.Param("id");

	var days []models.Day
	var timeEntries []models.TimeEntry


	const dayQuery = 
	`
		SELECT * 
		FROM days
		WHERE user_id = $1
	`

	const teQuery = 
	`
		SELECT * 
		FROM time_entries
		WHERE date = $1 AND user_id = $2
	`

	if 	err := db.DB.Select(&days,dayQuery,userId); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Days error": err.Error()})
		return
	}


	for i := 0; i < len(days); i++ {

		if 	err := db.DB.Select(&timeEntries,teQuery,days[i].Date,days[i].UserID); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"Time Entry error": err.Error()})
			return
		}

		days[i].TimeEntries = timeEntries

	}

	context.JSON(http.StatusOK, days)


}

func getAllDays(context *gin.Context){

	var days []models.Day
	const query = 
	`
		SELECT * 
		FROM days
		LIMIT 100
	`


	err := db.DB.Select(&days,query)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Days error": err.Error()})
		return
	}


	context.JSON(http.StatusOK, days)


}
