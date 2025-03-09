package routes

import (
	"demoProject/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func registerForEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}

	userId := c.GetInt64("userId")
	event, err := models.GetEventById(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get the event"})
		return
	}
	if event.UserId == userId {
		c.JSON(http.StatusBadRequest, gin.H{"message": "You cannot register for your own event"})
		return
	}
	err = event.RegisterForEvent(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register for the event"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Registered for the event"})
}

func unregisterFromEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}

	userId := c.GetInt64("userId")
	event, err := models.GetEventById(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get the event"})
		return
	}
	err = event.UnregisterFromEvent(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not unregister from the event"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Unregistered from the event"})
}
