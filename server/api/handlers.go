package handlers

import (
	"net/http"
	"server/database"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	database.CreateDummyUser()
	c.JSON(http.StatusAccepted, gin.H{"text": "hello"})
}

func GetUsers(c *gin.Context) {
	users := database.FindAllUsers()
	c.JSON(http.StatusFound, &users)
}
