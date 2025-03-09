package routes

import (
	"demoProject/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEvents(c *gin.Context) {
	events, err := models.GetEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get the events"})
		return
	}
	c.JSON(http.StatusOK, events)
}

func getEventById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}
	event, err := models.GetEventById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get the event"})
		return
	}
	c.JSON(http.StatusOK, event)
}

func createEvent(c *gin.Context) {
	var event models.Event
	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the data", "error": err.Error()})
		return
	}
	userId := c.GetInt64("userId")
	event.UserId = userId
	err = event.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save the event"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})
}

func updateEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}
	userid := c.GetInt64("userId")
	event, err := models.GetEventById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get the event"})
		return
	}
	if event.UserId != userid {
		c.JSON(http.StatusForbidden, gin.H{"message": "You are not allowed to update this event"})
		return
	}

	var updatedEvent models.Event
	err = c.ShouldBindJSON(&updatedEvent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the data", "error": err.Error()})
		return
	}
	updatedEvent.ID = event.ID
	err = updatedEvent.Upadate()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update the event"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Event updated successfully"})
}

func deleteEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}
	userid := c.GetInt64("userId")
	event, err := models.GetEventById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get the event"})
		return
	}
	if event.UserId != userid {
		c.JSON(http.StatusForbidden, gin.H{"message": "You are not allowed to delete this event"})
		return
	}
	err = event.DeleteEvent()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete the event"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}
