package queries

import (
	"errors"
	"fmt"
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

	const getTimeEntry = 
		`
		SELECT * 
		FROM time_entries
		WHERE id = $1
		`

	const query = 
		`
		UPDATE time_entries
		SET activity_id = $1, date = $2, start_time = $3, end_time = $4, total_hours = $5
		WHERE id = $6
		`

	date,startTime,endTime,total_hours,parseErr := tools.DateTimeHoursFormatter(request.Date,request.StartTime,request.EndTime)
	if parseErr != nil {
		e := errors.New("Failed to parse data " + parseErr.Error())
		return e,http.StatusBadRequest
	}

	var oldTimeEntry models.TimeEntry
	db.DB.Get(&oldTimeEntry,getTimeEntry,request.ActivityID, date, startTime, endTime,total_hours,request.ID)

	if(oldTimeEntry.Date != request.Date){
	}



	//Oppdaterer TimeEntry
	_,err := db.DB.Exec(query,request.ActivityID, date, startTime, endTime,total_hours,request.ID)

	if err != nil {
		e := errors.New("Failed to update Time Entry: " + err.Error())
		return e , http.StatusBadRequest
	}


	// TODO Update the day where timeEntry resides

	return nil ,  http.StatusOK
}


func DeleteDay(id int64) (error,int){

	//Delete day and get the date and user_id for the dates
    const dayQuery = `
        DELETE FROM days
        WHERE id = $1
        RETURNING date, user_id
    `
	var date string
	var user_id int
	err := db.DB.QueryRow(dayQuery,id).Scan(&date, &user_id)


	fmt.Printf("ID: %v", id)

    const teQuery = `
        DELETE FROM time_entries
        WHERE date = $1 AND user_id = $2
    `

	//Delete corresponding time_entries

	_,err2 := db.DB.Exec(teQuery,date,user_id);

	if err != nil {
		e := errors.New("Failed to delete day: " + err.Error())
		return e, http.StatusBadRequest
	}

	if err2 != nil {
		e := errors.New("Failed to delete time_entries: " + err.Error())
		return e, http.StatusBadRequest
	}



	return nil, http.StatusOK


}

func DeleteTimeEntry(id int64) (error,int){

	const teQuery = 
		`
		DELETE FROM time_entries
		WHERE id = $1
		RETURNING  date, user_id, total_hours
		`
	var date string
	var user_id int
	var total_hours float64

	err := db.DB.QueryRow(teQuery, id).Scan(&date, &user_id, &total_hours)
	if err != nil {
	}



	const dayQuery = `
		UPDATE days
		SET total_hours = total_hours - $1
		WHERE user_id = $2 AND date = $3
		RETURNING total_hours, id
		`

	var updatedTotal float64
	var day_id float64
	err2 := db.DB.QueryRow(dayQuery, total_hours, user_id, date).Scan(&updatedTotal, &day_id)


	if updatedTotal < 0.5{

	const dayQuery =
			`
			DELETE FROM days
			WHERE id = $1
			`

		_,err3 := db.DB.Exec(dayQuery,day_id)

		if err3 != nil {

		e := errors.New("Failed to delete day for that time_entry : " + err.Error())
		return e, http.StatusBadRequest

		}



	}





	if err != nil {
		e := errors.New("Failed to delete timeEntry : " + err.Error())
		return e, http.StatusBadRequest
	}

	if err2 != nil {
		e := errors.New("Failed to update Hours" + err.Error())
		return e, http.StatusBadRequest
	}
	return nil, http.StatusOK


}
