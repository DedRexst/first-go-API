package routes

import (
	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func deleteRegister(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	userId := context.GetInt64("userID")
	var event models.Event
	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Could not parse event id"})
		return
	}
	event.ID = eventId
	err = event.CancelRegistration(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel registration."})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Canceled!!!"})
}

func postRegister(context *gin.Context) {
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
	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register event."})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Successfully registered event."})
}
