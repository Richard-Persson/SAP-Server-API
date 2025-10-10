package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Richard-Persson/SAP-Server-API/db"
	"github.com/Richard-Persson/SAP-Server-API/internal/models"
	"github.com/Richard-Persson/SAP-Server-API/internal/payload/requests"
	"github.com/Richard-Persson/SAP-Server-API/internal/tools"
	"github.com/gin-gonic/gin"
)


func saveTimeEntry (context *gin.Context){

	var request requests.TimeEntryRequest

	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	const query = `
		INSERT INTO time_entries (user_id, activity_id, date, start_time, end_time, total_hours)
		VALUES ($1, $2, $3, $4, $5, $6) 
		RETURNING id, user_id, activity_id, date, start_time, end_time, total_hours
		`


	var time_entry models.TimeEntry

	// TODO The date format dd-mm-yyyy throws a panic in the program fix this
	date, dateParseErr := time.Parse("2006-01-02",request.Date)
	startTime, timeParseErr1 := time.Parse("15:04",request.StartTime)
	endTime, timeParseErr2 := time.Parse("15:04",request.EndTime)

	if dateParseErr != nil {
		context.JSON(http.StatusInternalServerError, dateParseErr.Error())
		return
	}

	if timeParseErr1 != nil {
		context.JSON(http.StatusInternalServerError, timeParseErr1.Error())
		return
	}
	if timeParseErr2 != nil {
		context.JSON(http.StatusInternalServerError, timeParseErr2.Error())
		return
	}

	total_hours := endTime.Sub(startTime).Hours()

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
	//TODO Make this a own function in tools package
	date, _ := time.Parse("2006-01-02",request.Date)
	startTime, _ := time.Parse("15:04",request.StartTime)
	endTime,_  := time.Parse("15:04",request.EndTime)

	total_hours := endTime.Sub(startTime).Hours()
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


