package controllers

import (
	"net/http"

	"github.com/KUNSALISA/Meeting-Room-Booking-System/config"
	"github.com/KUNSALISA/Meeting-Room-Booking-System/entity"
	"github.com/gin-gonic/gin"
)

func PostBooking (c *gin.Context) {
	var inputbook []entity.Booking

	if err := c.ShouldBindJSON(&inputbook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "error": err.Error()})
		return
	}

	for i := range inputbook {
		if inputbook[i].EndTime.Before((inputbook[i].StartTime)) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "EndTime ต้องงงมากกว่า StartTime",
				"intex": i,
			})
			return
		}
	}

	if err := config.DB().Create(&inputbook).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ไม่สามารถบันทึกได้",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": inputbook,
	})
}