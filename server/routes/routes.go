package routes

import (
	"github.com/dandimuzaki/badminton-server/handler"
	"github.com/dandimuzaki/badminton-server/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "pong",
    })
  })

	r.POST("/api/auth/register", handler.Register)
	r.POST("/api/auth/login", handler.Login)

	r.POST("/api/courts", handler.CreateCourt)
	r.GET("/api/courts", handler.GetCourts)
	r.GET("/api/courts/:id", handler.GetCourtByID)
	r.PUT("/api/courts/:id", handler.UpdateCourt)
	r.DELETE("/api/courts/:id", handler.DeleteCourt)

	r.GET("/api/timeslots", handler.GetTimeslots)
	r.GET("/api/timeslots/:id", handler.GetTimeslotByID)
	r.POST("/api/timeslots", handler.CreateTimeslot)
	r.DELETE("/api/timeslots/:id", handler.DeleteTimeslot)

	r.GET("/api/available-timeslots", handler.GetAvailableTimeslots)
	r.GET("/api/available-courts", handler.GetAvailableCourts)

	r.GET("/api/reservations", middleware.AuthMiddleware(), handler.GetUserReservations)
	r.POST("/api/reservations", middleware.AuthMiddleware(), handler.CreateReservation)
	r.PUT("/api/reservations/:id", middleware.AuthMiddleware(), handler.CancelReservation)

	r.POST("/api/payments/create", middleware.AuthMiddleware(), handler.CreatePayment)
	r.POST("/api/payments/notification", handler.PaymentNotification)
}
