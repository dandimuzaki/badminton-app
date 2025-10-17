package controllers

import (
	"net/http"
	"time"

	"github.com/dandimuzaki/badminton-server/initializers"
	"github.com/dandimuzaki/badminton-server/models"
	"github.com/gin-gonic/gin"
)

func CreateReservation(c *gin.Context) {
    var body struct {
        CourtID    uint   `json:"courtId"`
        Date       string `json:"date"`
        TimeSlotID uint   `json:"timeSlotId"`
    }

    userId, _ := c.Get("user_id") // from auth middleware

    if err := c.ShouldBindJSON(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Parse date safely
    resDate, err := time.Parse("2006-01-02", body.Date)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM-DD"})
        return
    }

    // Check if timeslot already booked
    var existing models.Reservation
    if err := initializers.DB.
        Where("court_id = ? AND date = ? AND time_slot_id = ?", body.CourtID, resDate, body.TimeSlotID).
        First(&existing).Error; err == nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Timeslot already booked"})
        return
    }

    // Create reservation
    reservation := models.Reservation{
        UserID:     userId.(uint),
        CourtID:    body.CourtID,
        Date:       resDate,
        TimeSlotID: body.TimeSlotID,
        Status:     "pending",
    }

    if err := initializers.DB.Create(&reservation).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create reservation"})
        return
    }

    // Mock payment
    payment := models.Payment{
        ReservationID: reservation.ID,
        Amount:        50000.00,
        Status:        "pending",
    }

    if err := initializers.DB.Create(&payment).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create payment"})
        return
    }

    // Link payment to reservation
    reservation.PaymentID = &payment.ID
    initializers.DB.Save(&reservation)

    c.JSON(http.StatusOK, gin.H{
        "message":     "Reservation created successfully",
        "reservation": reservation,
        "payment":     payment,
    })
}

func GetUserReservations(c *gin.Context) {
    // 1️⃣ Get user ID from JWT or session
    userID := c.GetUint("user_id") // depends on your middleware

    // 2️⃣ Fetch reservations + related court, timeslot, payment
    var reservations []models.Reservation
    if err := initializers.DB.
        Where("user_id = ?", userID).
        Preload("Court").
        Preload("Timeslot").
        Preload("Payment").
        Find(&reservations).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch reservations"})
        return
    }

    // 3️⃣ Return data
    c.JSON(http.StatusOK, gin.H{
        "message": "Fetched user reservations successfully",
        "data":    reservations,
    })
}

