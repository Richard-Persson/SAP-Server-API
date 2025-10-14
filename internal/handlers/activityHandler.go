package handlers

import (
	"net/http"
	"strconv"

	"github.com/Richard-Persson/SAP-Server-API/db"
	"github.com/Richard-Persson/SAP-Server-API/internal/models"
	"github.com/gin-gonic/gin"
)


func getActivityById(context *gin.Context)  {

	var id,_ = strconv.ParseInt(context.Param("id"),0,64)
	var activity models.Activity

	const query = 
		`
		SELECT * 
		FROM activities
		WHERE id = $1
		`


	err := db.DB.Get(&activity,query,id)
		
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	context.JSON(http.StatusOK, activity)
}


func getAllActivities(context *gin.Context)  {

	var activites []models.Activity

	const query = 
		`
		SELECT * 
		FROM activities
		`
	err := db.DB.Select(&activites,query)
		
	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	context.JSON(http.StatusOK, activites)
}


func getSingleDayByDayId(context *gin.Context){

	dayId := context.Param("id");

	var day models.Day
	var timeEntries []models.TimeEntry

	const dayQuery = 
		`
		SELECT * 
		FROM days
		WHERE id = $1
		`

	const teQuery = 
		`
		SELECT * 
		FROM time_entries
		WHERE date = $1 AND user_id = $2
		`
	err := db.DB.Get(&day, dayQuery, dayId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Days error": err.Error()})
		return
	}

	if 	err := db.DB.Select(&timeEntries,teQuery,day.Date,day.UserID); err != nil {

		context.JSON(http.StatusInternalServerError, gin.H{"Time Entry error": err.Error()})
		return
	}

	day.TimeEntries = timeEntries

	context.JSON(http.StatusOK, day)
}
