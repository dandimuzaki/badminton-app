package main

import (
	"log"
	"os"
	"time"

	"github.com/dandimuzaki/badminton-server/cmd"
	"github.com/dandimuzaki/badminton-server/database"
	"github.com/dandimuzaki/badminton-server/initializers"
	"github.com/dandimuzaki/badminton-server/routes"
	"github.com/dandimuzaki/badminton-server/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
		config, err := utils.ReadConfiguration()
	if err != nil {
		log.Fatalf("failed to read file config: %v", err)
	}

	db, err := database.InitDB(config.DB)
	if err != nil {
		log.Fatalf("failed to connect to postgres database: %v", err)
	}

	logger, err := utils.InitLogger(config.PathLogging, config.Debug)

	// migration
	err = database.AutoMigrate(db)
	if err != nil {
		log.Println(err)
	}

	route := routes.Wiring(db, logger, config)
	cmd.APiserver(route)

	// cron scheduler
	// route.Scheduler.Start()

	r := gin.Default()

	client := os.Getenv("CLIENT_HOST")

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", client},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	routes.SetupRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
