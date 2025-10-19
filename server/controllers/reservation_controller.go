package controllers

import (
	"net/http"
	"time"

	"github.com/dandimuzaki/badminton-server/initializers"
	"github.com/dandimuzaki/badminton-server/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateReservation creates a new reservation for a user
func CreateReservation(c *gin.Context) {
	var body struct {
		CourtID    uint   `json:"courtId" binding:"required"`
		Date       string `json:"date" binding:"required"`       // format: YYYY-MM-DD
		TimeSlotID uint   `json:"timeSlotId" binding:"required"` // FK to timeslot table
	}

	userID, _ := c.Get("user_id")

	// Validate input
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	// Parse date safely
	resDate, err := time.Parse("2006-01-02", body.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM-DD"})
		return
	}

	// Check if timeslot already booked for that court
	var existing models.Reservation
	if err := initializers.DB.
		Where("court_id = ? AND date = ? AND time_slot_id = ?", body.CourtID, resDate, body.TimeSlotID).
		First(&existing).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "This timeslot is already booked"})
		return
	}

	// Create reservation with "pending" status
	reservation := models.Reservation{
		UserID:     userID.(uint),
		CourtID:    body.CourtID,
		Date:       resDate,
		TimeSlotID: body.TimeSlotID,
		Status:     "pending",
	}

	if err := initializers.DB.Create(&reservation).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create reservation"})
		return
	}

	var fullReservation models.Reservation
	initializers.DB.Preload("User").Preload("Court").Preload("Timeslot").Preload("Payment").
		First(&fullReservation, reservation.ID)

	c.JSON(http.StatusOK, gin.H{
		"message": "Reservation created successfully",
		"data":    fullReservation,
	})
}

// GetUserReservations returns all reservations for logged-in user
func GetUserReservations(c *gin.Context) {
	userID := c.GetUint("user_id")

	var reservations []models.Reservation
	if err := initializers.DB.
		Where("user_id = ?", userID).
		Preload("User").
		Preload("Court").
		Preload("Timeslot").
		Preload("Payment").
		Order("date DESC").
		Find(&reservations).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch reservations"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Fetched user reservations successfully",
		"data":    reservations,
	})
}

// CancelReservation allows user to cancel a pending reservation
func CancelReservation(c *gin.Context) {
	id := c.Param("id")
	userID := c.GetUint("user_id")

	var reservation models.Reservation
	if err := initializers.DB.First(&reservation, "id = ? AND user_id = ?", id, userID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Reservation not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch reservation"})
		}
		return
	}

	if reservation.Status != "pending" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Only pending reservations can be canceled"})
		return
	}

	reservation.Status = "canceled"
	initializers.DB.Save(&reservation)

	c.JSON(http.StatusOK, gin.H{
		"message": "Reservation canceled successfully",
	})
}

// AdminGetAllReservations (optional) – admin dashboard
func AdminGetAllReservations(c *gin.Context) {
	var reservations []models.Reservation
	if err := initializers.DB.
		Preload("User").
		Preload("Court").
		Preload("Timeslot").
		Order("date DESC").
		Find(&reservations).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch reservations"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Fetched all reservations successfully",
		"data":    reservations,
	})
}
