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
    AllowMethods:     []string{"GET", "POST", "OPTIONS"},
    AllowHeaders:     []string{"Content-Type"},
    AllowCredentials: true,
    MaxAge:           12 * time.Hour,
  }))


	router.GET("/users", listUsers)
	router.GET("/users/:id", getUser)


	router.GET("/months",listMonths)

	router.POST("/login", login)

	router.POST("/register", register)


  return router
}
