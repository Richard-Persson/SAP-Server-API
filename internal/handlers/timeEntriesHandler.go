package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Richard-Persson/SAP-Server-API/db"
	"github.com/Richard-Persson/SAP-Server-API/internal/models"
	"github.com/Richard-Persson/SAP-Server-API/internal/payload/requests"
	"github.com/Richard-Persson/SAP-Server-API/internal/queries"
	"github.com/Richard-Persson/SAP-Server-API/internal/tools"
	"github.com/gin-gonic/gin"
)


func createTimeEntry (context *gin.Context){

	var request requests.TimeEntryRequest
	var time_entry models.TimeEntry
	var day models.Day

	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	const query = `
		INSERT INTO time_entries (user_id, activity_id, date, start_time, end_time, total_hours)
		VALUES ($1, $2, $3, $4, $5, $6) 
		RETURNING id, user_id, activity_id, date, start_time, end_time, total_hours
		`

	date,startTime,endTime,total_hours_entries,parseErr := tools.DateTimeHoursFormatter(request.Date,request.StartTime,request.EndTime)

	if parseErr != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"parse error": parseErr.Error()})
		return
	}


	//Insert time entry in DB
	err :=	db.DB.Get(&time_entry,query,request.UserID,request.ActivityID,date,startTime, endTime, total_hours_entries)

	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}



	//===============================================================================================
	//Create a new day in the DB with the same date as the timeEntry if that day dosen't already exist
	//TODO refactor this later?

	const dayQuery =
		`
		SELECT * 
		FROM days
		WHERE date = $1 AND user_id = $2
		`
	dayErr := db.DB.Get(&day,dayQuery,date,request.UserID)


	if dayErr != nil {
		createDay(request.UserID, date, total_hours_entries)
	}else {

		//If day already exists
		err := updateDay(date,total_hours_entries,request.UserID)

		if  err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Error ": err.Error()})
			return 
		}


	}

	context.JSON(http.StatusCreated, time_entry)

}


func getAllTimeEntries(context *gin.Context) {

	var timeEntries []models.TimeEntry
	user_id, err := strconv.ParseInt(context.Param("id"),0,64)

	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	queries.GetTimeEntriesByUserId(&timeEntries,user_id,context)
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



func createDay(user_id int64, date time.Time,total_hours float64 ) error{

	var day models.Day


	const query = `
		INSERT INTO days(date, user_id, total_hours) 
		VALUES ($1, $2, $3)  
		`
	fmt.Println("======================================")
	fmt.Printf("HOURS ALEREADY IN DATE: %v", total_hours)
	fmt.Println("======================================")

	err:= db.DB.Get(&day,query,date,user_id,total_hours);

	if err != nil { return err }

	return nil

}


func updateDay (date time.Time, total_hours_entries float64, user_id int64) error {


	//Get current hours
	const getQuery = 
		`
		SELECT total_hours 
		FROM days
		WHERE date = $1 and user_id = $2
		`

	var current_hours float64
	getErr := db.DB.Get(&current_hours,getQuery,date,user_id);
	if getErr != nil { return getErr }

	var new_hours = current_hours + total_hours_entries

	if(new_hours > 24){
		error := errors.New("Cant have more than 24 hours in a day")
		return error
	}

	//Add new hours
	const insertQuery = 
		`
		UPDATE days
		SET total_hours = $1
		WHERE date = $2 and user_id = $3
		`


	_,insertErr := db.DB.Exec(insertQuery,current_hours + total_hours_entries,date,user_id)
	if insertErr != nil { return insertErr }



	return nil

}

