package main

import (
	"log"

	"github.com/dandimuzaki/badminton-server/cmd"
	"github.com/dandimuzaki/badminton-server/database"
	"github.com/dandimuzaki/badminton-server/routes"
	"github.com/dandimuzaki/badminton-server/utils"
)

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
}
