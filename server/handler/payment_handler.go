package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dandimuzaki/badminton-server/initializers"
	"github.com/dandimuzaki/badminton-server/model"
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

	var reservation model.Reservation
	if err := initializers.DB.Preload("Court").First(&reservation, "id = ?", body.ReservationID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reservation not found"})
		return
	}
	if reservation.UserID != userId {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized access to this reservation"})
		return
	}

	var user model.User
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
	payment := model.Payment{
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

	orderID := payload["order_id"].(string)
	status := payload["transaction_status"].(string)

	// Update payment status
	initializers.DB.Model(&model.Payment{}).
		Where("transaction_id = ?", orderID). // your DB stores orderID here
		Update("status", status)

	// Update reservation
	if status == "settlement" || status == "capture" {
		initializers.DB.Exec(`
			UPDATE reservations 
			SET status = 'paid'
			WHERE id = (
				SELECT reservation_id FROM payments WHERE transaction_id = ?
			)
		`, orderID)
	}

	if status == "cancel" || status == "expire" || status == "deny" {
		initializers.DB.Exec(`
			UPDATE reservations 
			SET status = 'failed'
			WHERE id = (
				SELECT reservation_id FROM payments WHERE transaction_id = ?
			)
		`, orderID)
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
