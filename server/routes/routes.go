package routes

import (
	"sync"
	"time"

	"github.com/dandimuzaki/badminton-server/adaptor"
	"github.com/dandimuzaki/badminton-server/handler"
	"github.com/dandimuzaki/badminton-server/infra"
	"github.com/dandimuzaki/badminton-server/middleware"
	"github.com/dandimuzaki/badminton-server/repository"
	"github.com/dandimuzaki/badminton-server/usecase"
	"github.com/dandimuzaki/badminton-server/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type App struct {
	Route  *gin.Engine
	Stop   chan struct{}
	WG     *sync.WaitGroup
	Config utils.Configuration
}

func Wiring(db *gorm.DB, log *zap.Logger, config utils.Configuration) *App {
	r := gin.Default()
	
	r.Use(cors.New(cors.Config{
    AllowOrigins: []string{"http://localhost:3000", config.ClientHost},
    AllowMethods: []string{
        "GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS",
    },
    AllowHeaders: []string{
        "Origin", "Content-Type", "Authorization", "Accept",
    },
    AllowCredentials: true,
    MaxAge: 12 * time.Hour,
	}))

	r1 := r.Group("/api/v1")

	stop := make(chan struct{})
	wg := &sync.WaitGroup{}

	tx := infra.NewGormTxManager(db)
	// emailService := infra.NewEmailService(config.SMTP, log)
	tokenService := infra.NewTokenService(config.JWTSecret, config.Issuer)

	repo := repository.NewRepository(db, log)
	usecase := usecase.NewUsecase(tx, repo, tokenService, log, config)
	handler := adaptor.NewHandler(usecase, log)

	ApiV1(r1, &handler)

	return &App{
		Route:  r,
		Stop:   stop,
		WG:     wg,
		Config: config,
	}
}

func ApiV1(r *gin.RouterGroup, handler *adaptor.Handler) {
	r.POST("/auth/register", handler.AuthHandler.Register)
	r.POST("/auth/login", handler.AuthHandler.Login)

	r.GET("/timeslots", handler.TimeslotHandler.GetAvailableTimeslots)
	r.GET("/courts", handler.CourtHandler.GetAvailableCourts)

	r.GET("/reservations", middleware.AuthMiddleware(), handler.ReservationHandler.GetReservationHistory)
	r.POST("/reservations", middleware.AuthMiddleware(), handler.ReservationHandler.CreateReservation)
	r.PUT("/reservations/:id/cancel", middleware.AuthMiddleware(), handler.ReservationHandler.CancelReservation)

	r.POST("/payments/create", middleware.AuthMiddleware(), handler.PaymentHandler.CreatePayment)
	r.POST("/payments/notification", handler.PaymentHandler.PaymentCallback)
}

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
