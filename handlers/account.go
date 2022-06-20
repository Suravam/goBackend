package handlers

/*
This module handles all the user account related request / response
*/

import (
	"net/http"

	"github.com/Suravam/goBackend/models"
	"github.com/gin-gonic/gin"
)

//HTTP POST Handler that adds users on signup
func addUser(c *gin.Context) {

	var newUser models.User

	//receive JSON
	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	//TODO: Handle Database Insert

	c.IndentedJSON(http.StatusCreated, newUser)
}
