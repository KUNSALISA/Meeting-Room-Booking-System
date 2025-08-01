package controllers

import (
	"net/http"

	"github.com/KUNSALISA/Meeting-Room-Booking-System/config"
	"github.com/KUNSALISA/Meeting-Room-Booking-System/entity"
	"github.com/gin-gonic/gin"
)

func GetAllRoom (c *gin.Context) {
	var room []entity.Room

	config.DB().Preload("Type").Preload("Status").Find(&room)

	c.JSON(http.StatusOK, gin.H{
		"data": room,
	})
}