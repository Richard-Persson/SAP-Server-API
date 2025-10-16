package queries

import (
	"errors"
	"net/http"

	"github.com/Richard-Persson/SAP-Server-API/db"
	"github.com/Richard-Persson/SAP-Server-API/internal/models"
	"github.com/Richard-Persson/SAP-Server-API/internal/payload/requests"
	"github.com/Richard-Persson/SAP-Server-API/internal/tools"
)

/*
Gets all entries for a user given their ID
*/
func GetTimeEntriesByUserId(timeEntries *[]models.TimeEntry ,user_id int64) (error, int)  {

	const timeEntriesQuery= `
		SELECT * 
		FROM time_entries 
		WHERE user_id = $1
		`
	
	error := db.DB.Select(timeEntries,timeEntriesQuery,user_id)

	if error != nil {
		e := errors.New("Could not find time entries given user_id")
		return e,http.StatusBadRequest
	}


	//Removes T00:00:00Z From Date attribute
	tools.RemoveTZ(timeEntries)
	return nil, http.StatusOK
}


/*
Gets all entries for a user given their ID
*/
func UpdateTimeEntry(request requests.UpdateTimeEntryRequest, timeEntry *models.TimeEntry) (error, int) {


	const query = 
		`
		UPDATE time_entries
		SET activity_id = $1, date = $2, start_time = $3, end_time = $4, total_hours = $5
		WHERE id = $6
		`

	date,startTime,endTime,total_hours,parseErr := tools.DateTimeHoursFormatter(request.Date,request.StartTime,request.EndTime)
	if parseErr != nil {
		e := errors.New("Failed to parse data" + parseErr.Error())
		return e,http.StatusBadRequest
	}

	_,err := db.DB.Exec(query,request.ActivityID, date, startTime, endTime,total_hours,request.ID)

	if err != nil {
		e := errors.New("Failed to update Time Entry: " + err.Error())
		return e , http.StatusBadRequest
	}


	return nil ,  http.StatusOK

}

