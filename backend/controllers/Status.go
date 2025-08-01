package controllers

import (
	"github.com/KUNSALISA/Meeting-Room-Booking-System/config"
	"github.com/KUNSALISA/Meeting-Room-Booking-System/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetStatus(c *gin.Context) {
	var status entity.Status

	config.DB().Find(&status)

	c.JSON(http.StatusOK, gin.H{
		"data": status,
	})
}
