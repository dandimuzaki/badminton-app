package controllers

import (
	"net/http"
	"time"

	"github.com/dandimuzaki/badminton-server/initializers"
	"github.com/dandimuzaki/badminton-server/models"
	"github.com/gin-gonic/gin"
)

func CreateTimeslot(c *gin.Context) {
	var body struct {
		StartTime string `json:"startTime"`
		EndTime   string `json:"endTime"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Parse input string (e.g. "08:00") into time.Time
	start, err := time.Parse("15:04", body.StartTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid startTime format, use HH:MM"})
		return
	}

	end, err := time.Parse("15:04", body.EndTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid endTime format, use HH:MM"})
		return
	}

	timeslot := models.Timeslot{
		StartTime: start,
		EndTime:   end,
	}

	if err := initializers.DB.Create(&timeslot).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Timeslot created successfully",
		"data":    timeslot,
	})
}


func GetTimeslots(c *gin.Context) {
	var timeslots []models.Timeslot
	initializers.DB.Find(&timeslots)
	c.JSON(http.StatusOK, gin.H{"data": timeslots})
}

func GetTimeslotByID(c *gin.Context) {
	id := c.Param("id")
	var timeslot models.Timeslot
	if err := initializers.DB.First(&timeslot, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Timeslot not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": timeslot})
}

func DeleteTimeslot(c *gin.Context) {
	id := c.Param("id")

	// Check if court exists first
	var timeslot models.Timeslot
	if err := initializers.DB.First(&timeslot, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Timeslot not found"})
		return
	}

	// Delete the record
	if err := initializers.DB.Delete(&timeslot).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Timeslot deleted successfully",
	})
}
