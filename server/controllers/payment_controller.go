package controllers

import (
	"net/http"

	"github.com/dandimuzaki/badminton-server/initializers"
	"github.com/dandimuzaki/badminton-server/models"
	"github.com/gin-gonic/gin"
)

func CreatePayment(c *gin.Context) {
    var body struct {
        ReservationID uint   `json:"reservationId"`
        Amount        float64 `json:"amount"`
    }

    if err := c.ShouldBindJSON(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Normally you'd call payment gateway API here (e.g. Midtrans)
    // but for the technical test you can mock this step.

    payment := models.Payment{
        ReservationID: body.ReservationID,
        Amount:        body.Amount,
        Status:        "pending",
    }

    if err := initializers.DB.Create(&payment).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create payment"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Payment created successfully",
        "data":    payment,
    })
}

// webhook/notification handler
func PaymentNotification(c *gin.Context) {
    var payload struct {
        TransactionID string `json:"transaction_id"`
        Status        string `json:"status"`
    }

    if err := c.ShouldBindJSON(&payload); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Find payment by transaction ID
    var payment models.Payment
    if err := initializers.DB.Where("transaction_id = ?", payload.TransactionID).First(&payment).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
        return
    }

    // Update payment status
    initializers.DB.Model(&payment).Update("status", payload.Status)

    // If payment succeeded, mark reservation as paid
    if payload.Status == "success" {
        initializers.DB.Model(&models.Reservation{}).
            Where("id = ?", payment.ReservationID).
            Update("status", "paid")
    }

    c.JSON(http.StatusOK, gin.H{"message": "Payment notification processed"})
}

