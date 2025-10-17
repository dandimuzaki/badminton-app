package main

import (
	"os"

	"github.com/dandimuzaki/badminton-server/initializers"
	"github.com/dandimuzaki/badminton-server/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	routes.SetupRoutes(r)

	port := os.Getenv(("PORT"))
	r.Run(":" + port)
}
