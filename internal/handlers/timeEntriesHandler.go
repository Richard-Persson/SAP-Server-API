package handlers

import (
	"net/http"
	"strconv"

	"github.com/Richard-Persson/SAP-Server-API/db"
	"github.com/Richard-Persson/SAP-Server-API/internal/models"
	"github.com/Richard-Persson/SAP-Server-API/internal/payload/requests"
	"github.com/Richard-Persson/SAP-Server-API/internal/tools"
	"github.com/gin-gonic/gin"
)


func saveTimeEntry (context *gin.Context){

	var request requests.TimeEntryRequest
	var time_entry models.TimeEntry

	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	const query = `
		INSERT INTO time_entries (user_id, activity_id, date, start_time, end_time, total_hours)
		VALUES ($1, $2, $3, $4, $5, $6) 
		RETURNING id, user_id, activity_id, date, start_time, end_time, total_hours
		`

	date,startTime,endTime,total_hours,parseErr := tools.DateTimeHoursFormatter(request.Date,request.StartTime,request.EndTime)

	if parseErr != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"parse error": parseErr.Error()})
		return
	}


	err :=	db.DB.Get(&time_entry,query,request.UserID,request.ActivityID,date,startTime, endTime, total_hours)

	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	context.JSON(http.StatusCreated, time_entry)

}


func getAllTimeEntries(context *gin.Context) {

	id, err := strconv.ParseInt(context.Param("id"),0,64)

	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}


	const query = `
		SELECT * 
		FROM time_entries 
		WHERE user_id = $1
		`
	var timeEntries []models.TimeEntry 

	dbErr := db.DB.Select(&timeEntries,query,id)

	//Removes T00:00:00Z From Date attribute
	tools.RemoveTZ(&timeEntries)



	if dbErr != nil {
		context.JSON(http.StatusInternalServerError, dbErr.Error())
		return
	}
	context.JSON(http.StatusAccepted, timeEntries)
}



func updateTimeEntry(context *gin.Context){


	var request requests.UpdateTimeEntryRequest
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	const query = 
		`
		UPDATE time_entries
		SET activity_id = $1, date = $2, start_time = $3, end_time = $4, total_hours = $5
		WHERE id = $6
		`

	date,startTime,endTime,total_hours,parseErr := tools.DateTimeHoursFormatter(request.Date,request.StartTime,request.EndTime)
	if parseErr != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"parse error": parseErr.Error()})
		return
	}
	var timeEntry models.TimeEntry

	//Update timeEntry
	_,err := db.DB.Exec(query,request.ActivityID, date, startTime, endTime,total_hours,request.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not update timeEntry",
			"errorMessage": err.Error(),
		})
		return
	}

	//Get the updated entry
	db.DB.Get(&timeEntry,"SELECT * FROM time_entries WHERE id = $1",request.ID)
	tools.RemoveSingleTZ(&timeEntry)
	context.JSON(http.StatusOK, timeEntry)

}


