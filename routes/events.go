package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/qsmsoft/go-rest-api/models"
)

func getEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve events."})
		return
	}
	c.JSON(http.StatusOK, events)
}

func getEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID."})
		return
	}
	event, err := models.GetEventByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve event."})
		return
	}

	c.JSON(http.StatusOK, event)
}

func createEvent(c *gin.Context) {
	var event models.Event
	err := c.ShouldBindJSON(&event)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	userId := c.GetInt64("userId")
	event.UserID = userId

	err = event.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create event."})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

func updateEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID."})
		return
	}

	userId := c.GetInt64("userId")
	event, err := models.GetEventByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	if event.UserID != userId {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "You are not authorized to update this event."})
		return
	}

	var updatedEvent models.Event
	err = c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update event."})
		return
	}

	updatedEvent.ID = id
	err = event.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update event."})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event updated!", "event": updatedEvent})
}

func deleteEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID."})
		return
	}
	userId := c.GetInt64("userId")
	event, err := models.GetEventByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	if event.UserID != userId {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "You are not authorized to delete this event."})
		return
	}

	err = event.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete event."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event deleted!"})
}
