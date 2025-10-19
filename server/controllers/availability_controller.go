package controllers

import (
	"net/http"

	"github.com/dandimuzaki/badminton-server/initializers"
	"github.com/dandimuzaki/badminton-server/models"
	"github.com/gin-gonic/gin"
)

func GetAvailableCourts(c *gin.Context) {
	date := c.Query("date")
	timeslotID := c.Query("timeSlotId")

	var courts []models.Court

	// Use subquery to exclude booked courts
	subQuery := initializers.DB.Model(&models.Reservation{}).
		Select("court_id").
		Where("date = ? AND time_slot_id = ? AND status = ?", date, timeslotID, "pending")

	if err := initializers.DB.
		Where("id NOT IN (?)", subQuery).
		Find(&courts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": courts})
}

func GetAvailableTimeslots(c *gin.Context) {
	date := c.Query("date")

	var timeslots []models.Timeslot

	// First, count total courts
	var totalCourts int64
	initializers.DB.Model(&models.Court{}).Count(&totalCourts)

	// Subquery: timeslot IDs that are fully booked
	subQuery := initializers.DB.
		Model(&models.Reservation{}).
		Select("time_slot_id").
		Where("date = ? AND status = ?", date, "pending").
		Group("time_slot_id").
		Having("COUNT(court_id) >= ?", totalCourts)

	// Fetch timeslots that are not fully booked
	if err := initializers.DB.
		Where("id NOT IN (?)", subQuery).
		Find(&timeslots).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": timeslots})
}
