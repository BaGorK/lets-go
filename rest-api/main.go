package main

import (
	"net/http"

	"github.com/BaGorK/lets-go/rest-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/events", getEvents)
	router.POST("/events", createEvent)

	router.Run()
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event

	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event.Save()

	context.JSON(http.StatusCreated, gin.H{
		"message": "Event created successfully",
		"event":   event,
	})
}
