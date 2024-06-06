package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qsmsoft/go-rest-api/db"
	"github.com/qsmsoft/go-rest-api/models"
)

func main() {
	db.InitDB()
	defer func() {
		if err := db.DB.Close(); err != nil {
			log.Fatalf("Failed to close database: %v", err)
		}
	}()
	fmt.Println("Database initialized and tables created successfully.")

	server := gin.Default()
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080")
}

func getEvents(c *gin.Context) {
	events := models.GetAllEvents()
	c.JSON(http.StatusOK, events)
}

func createEvent(c *gin.Context) {
	var event models.Event
	err := c.ShouldBindJSON(&event)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	event.ID = 1
	event.UserID = 1

	event.Save()

	c.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}
