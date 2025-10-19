package controllers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dandimuzaki/badminton-server/initializers"
	"github.com/dandimuzaki/badminton-server/models"
	"github.com/gin-gonic/gin"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

func CreatePayment(c *gin.Context) {
	var body struct {
		ReservationID uint    `json:"reservationId"`
		Amount        float64 `json:"amount"`
	}

	userId := c.GetUint("user_id")

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	var reservation models.Reservation
	if err := initializers.DB.Preload("Court").First(&reservation, "id = ?", body.ReservationID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reservation not found"})
		return
	}
	if reservation.UserID != userId {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized access to this reservation"})
		return
	}

	var user models.User
	if err := initializers.DB.First(&user, userId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	serverKey := os.Getenv("MIDTRANS_SERVER_KEY")
	if serverKey == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "MIDTRANS_SERVER_KEY not configured"})
		return
	}

	// ✅ Initialize Snap client directly
	var s = snap.Client{}
	s.New(os.Getenv("MIDTRANS_SERVER_KEY"), midtrans.Sandbox)

	orderID := fmt.Sprintf("ORDER-%d-%d", reservation.ID, userId)

	// ✅ Use correct struct names for request
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderID,
			GrossAmt: int64(body.Amount),
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: user.Name,
			Email: user.Email,
		},
		Items: &[]midtrans.ItemDetails{
			{
				ID:    fmt.Sprintf("court-%d", reservation.CourtID),
				Price: int64(reservation.Court.Price),
				Qty:   1,
				Name:  reservation.Court.Name,
			},
		},
	}

	// ✅ Create transaction
	snapResp, err := s.CreateTransaction(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// ✅ Save payment record
	payment := models.Payment{
		UserID:        userId,
		ReservationID: reservation.ID,
		Amount:        body.Amount,
		Status:        "pending",
		TransactionID: orderID,
	}
	initializers.DB.Create(&payment)

	c.JSON(http.StatusOK, gin.H{
		"message":     "Payment created successfully",
		"snapToken":   snapResp.Token,
		"redirectUrl": snapResp.RedirectURL,
	})
}

func PaymentNotification(c *gin.Context) {
	var payload map[string]interface{}
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	orderID, _ := payload["order_id"].(string)
	status, _ := payload["transaction_status"].(string)

	switch status {
	case "settlement":
		initializers.DB.Model(&models.Payment{}).
			Where("transaction_id = ?", orderID).
			Update("status", "success")

		initializers.DB.Model(&models.Reservation{}).
			Joins("JOIN payments ON payments.reservation_id = reservations.id").
			Where("payments.transaction_id = ?", orderID).
			Update("reservations.status", "paid")

	case "cancel", "expire", "deny":
		initializers.DB.Model(&models.Payment{}).
			Where("transaction_id = ?", orderID).
			Update("status", status)
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
