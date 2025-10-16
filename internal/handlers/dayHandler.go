package handlers

import (
	"net/http"

	"github.com/Richard-Persson/SAP-Server-API/db"
	"github.com/Richard-Persson/SAP-Server-API/internal/models"
	"github.com/Richard-Persson/SAP-Server-API/internal/tools"
	"github.com/gin-gonic/gin"
)
func getAllDaysByUserId(context *gin.Context){

	userId := context.Param("id");
	var days []models.Day

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

	//Get all time entries for all days
	for i := 0; i < len(days); i++ {

	var timeEntries []models.TimeEntry //New list for each day

		if 	err := db.DB.Select(&timeEntries,teQuery,days[i].Date,days[i].UserID); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"Time Entry error": err.Error()})
			return
		}

		tools.RemoveTZ(&timeEntries)
		tools.DateFormatter(&days[i].Date)
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
