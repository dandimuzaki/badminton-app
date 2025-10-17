package routes

import (
	"github.com/dandimuzaki/badminton-server/controllers"
	"github.com/dandimuzaki/badminton-server/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "pong",
    })
  })

	r.POST("/api/register", controllers.Register)
	r.POST("/api/login", controllers.Login)

	auth := r.Group("/api")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/me", func(c *gin.Context) {
			userID, _ := c.Get("user_id")
			c.JSON(200, gin.H{"user_id": userID})
		})
	}

	r.POST("/api/courts", controllers.CreateCourt)
	r.GET("/api/courts", controllers.GetCourts)
	r.GET("/api/courts/:id", controllers.GetCourtByID)
	r.PUT("/api/courts/:id", controllers.UpdateCourt)
	r.DELETE("/api/courts/:id", controllers.DeleteCourt)

	r.GET("/api/timeslots", controllers.GetTimeslots)
	r.GET("/api/timeslots/:id", controllers.GetTimeslotByID)
	r.POST("/api/timeslots", controllers.CreateTimeslot)
	r.DELETE("/api/timeslots/:id", controllers.DeleteTimeslot)

	r.GET("/api/available-timeslots", controllers.GetAvailableTimeslots)
	r.GET("/api/available-courts", controllers.GetAvailableCourts)

	r.GET("/api/reservations", middleware.AuthMiddleware(), controllers.GetUserReservations)
	r.POST("/api/reservations", middleware.AuthMiddleware(), controllers.CreateReservation)

	r.POST("/api/payments/create", middleware.AuthMiddleware(), controllers.CreatePayment)
	r.POST("/api/payments/notification", middleware.AuthMiddleware(), controllers.PaymentNotification)
}
