package main

import (
	routes "chat-app-back/src/routes"
	apiRoute "chat-app-back/src/routes/api"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const addr string = "localhost:8080"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Setup routes
	router := gin.Default()

	// Api routes
	api := router.Group("/api")
	{
		apiRoute.MessageRoutes(api)
		apiRoute.ProfileRoutes(api)
	}

	// Authentication routes
	auth := router.Group("/auth")
	{
		routes.AuthenticationRoutes(auth)
	}

	// Run server
	router.Run(addr)
}
