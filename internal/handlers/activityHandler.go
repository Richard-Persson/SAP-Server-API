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

