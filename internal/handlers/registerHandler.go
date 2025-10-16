
package handlers

import(
	"fmt"
	"net/http"
	"github.com/Richard-Persson/SAP-Server-API/db"
	"github.com/Richard-Persson/SAP-Server-API/internal/tools"
	"github.com/Richard-Persson/SAP-Server-API/internal/payload/requests"
	"github.com/Richard-Persson/SAP-Server-API/internal/models"
	"github.com/gin-gonic/gin"
)


func register(c *gin.Context){

  var request requests.RegisterRequest 
  if err := c.ShouldBindJSON(&request); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }


  // Hash the plaintext password
  hash, err := tools.HashPassword(request.Password)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
    return
  }

  // Insert user into the database
	const query = 
		`
		INSERT INTO users (email, first_name, last_name, mobile, password, billing_code_id)
		VALUES ($1, $2, $3, $4, $5, 1)
		RETURNING id, first_name, last_name, email, mobile, password, billing_code_id
		`
  var user models.User
  err = db.DB.Get(&user, query, request.Email, request.FirstName, request.LastName, request.Mobile,hash)

  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create user"})
		fmt.Printf(" %v", err)
    return
  }

	//Returns the user which was created
	c.JSON(http.StatusCreated, user)
}
