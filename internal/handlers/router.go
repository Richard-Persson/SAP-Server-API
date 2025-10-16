package handlers

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)


func Router() http.Handler {
	router := gin.Default()


	// CORS config for React dev server
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))


	//Users
	router.GET("/users", listUsers)
	router.GET("/users/:id", getUserById)
	router.GET("/users/entries/:id", getAllTimeEntriesGivenUserId)

	//Billing codes
	router.GET("/billingcodes", getAllBillingCodes)
	router.GET("/billingcodes/:id",getBillingCodeById)

	//Activities
	router.GET("/activities", getAllActivities)
	router.GET("/activities/:id", getActivityById)


	router.POST("/timeEntry", createTimeEntry)
	router.PATCH("/timeEntry", updateTimeEntry)
	router.GET("/timeEntry/day/:id", getSingleDayWithTimeEntries)


	router.GET("days/all/:id", getAllDaysByUserId)

	router.GET("days/", getAllDays)

	router.POST("/login", login)
	router.POST("/register", register)


	return router
}
