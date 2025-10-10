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


	router.GET("/users", listUsers)
	router.GET("/users/:id", getUserById)
	router.GET("/users/entries/:id", getAllTimeEntries)

	router.GET("/billingcodes", getAllBillingCodes)
	router.GET("/billingcodes/:id",getBillingCodeById)

	router.GET("/activities", getAllActivities)
	router.GET("/activities/:id", getActivityById)

	router.POST("/timeEntry", saveTimeEntry)



	router.POST("/login", login)
	router.POST("/register", register)


  return router
}
