package queries

import (
	"net/http"

	"github.com/Richard-Persson/SAP-Server-API/db"
	"github.com/Richard-Persson/SAP-Server-API/internal/models"
	"github.com/Richard-Persson/SAP-Server-API/internal/tools"
	"github.com/gin-gonic/gin"
)



/*
Gets all entries for a user given their ID
*/
func GetTimeEntriesByUserId(timeEntries *[]models.TimeEntry ,user_id int64, context *gin.Context) ()  {



	const timeEntriesQuery= `
		SELECT * 
		FROM time_entries 
		WHERE user_id = $1
		`
	
	error := db.DB.Select(timeEntries,timeEntriesQuery,user_id)

	if error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Error": error.Error()})
		return
	}
	//Removes T00:00:00Z From Date attribute
	tools.RemoveTZ(timeEntries)

}
