package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Richard-Persson/SAP-Server-API/db"
	"github.com/Richard-Persson/SAP-Server-API/internal/models"
	"github.com/gin-gonic/gin"
)

func getBillingCodeById(context *gin.Context)  {

	var id,_ = strconv.ParseInt(context.Param("id"),0,64)
	var billingCode models.BillingCode
	fmt.Println(id)

	const query = 
		`
		SELECT * 
		FROM billing_codes
		WHERE id = $1
		`
		
	if err := db.DB.Get(&billingCode,query,id); err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	context.JSON(http.StatusOK, billingCode)
}

func getAllBillingCodes(context *gin.Context)  {

	var billingCodes []models.BillingCode
	const query = 
		`
		SELECT * 
		FROM billing_codes
		`
	if err := db.DB.Select(&billingCodes,query); err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	context.JSON(http.StatusOK, billingCodes)
}
