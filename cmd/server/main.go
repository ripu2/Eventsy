package main

import (
	"fmt"
	"os"

	"example.com/event-management/internal/db"
	"example.com/event-management/internal/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	db.InitDB()
	_ = godotenv.Load()
	server := gin.Default()
	routes.SetupRoutes(server)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("🚀 Server is up an running on http://localhost:8080")
	server.Run(":" + port) //localhost:8080

}
