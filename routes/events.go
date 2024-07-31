package routes

import (
	"example.com/rest-api/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func getEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Could not parse event id"})
	}

	event, err := models.GetEventById(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		fmt.Print(err)
		return
	}

	context.JSON(http.StatusOK, *event)
}

func postEvents(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	userId := context.GetInt64("userID")
	event.UserID = userId
	err = event.Save()
	if err != nil {
		panic(err)
	}
	context.JSON(http.StatusOK, event)
}

func getEvents(context *gin.Context) {
	events, err := models.GetEvents()
	if err != nil {
		return
	}
	context.JSON(http.StatusOK, events)
}

func putEvents(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Could not parse event id"})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}
	userId := context.GetInt64("userID")
	if event.UserID != userId {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "This event doesn't belong to you."})
		return
	}

	var upgradeEvent models.Event
	err = context.ShouldBindJSON(&upgradeEvent)

	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Could not parse request data."})
		return
	}

	upgradeEvent.ID = eventId
	err = upgradeEvent.UpdateEvent()
	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Could not save changed data."})
		return
	}

	context.JSON(http.StatusOK, upgradeEvent)
}

func deleteEvents(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Could not parse event id"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}
	userId := context.GetInt64("userID")
	if event.UserID != userId {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "This event doesn't belong to you."})
		return
	}

	models.DeleteEventById(eventId)
}
