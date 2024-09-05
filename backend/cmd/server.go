package main

import (
	"backend/api/route"
	"backend/internal/controller"
	"backend/internal/dao"
	"backend/internal/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	//
	db, err := dao.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Initialize services with proper arguments
	linkService := service.NewLinkService(dao.NewLinkDAO(db)) // Pass db instance
	userService := service.NewUserService(db)                 // Pass db instance

	// Initialize Gin router
	r := gin.Default()

	// CORS middleware configuration
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Allow frontend domain
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Initialize controllers
	linkController := controller.NewLinkController(linkService)
	userController := controller.NewUserController(userService)

	// Register routes with the controllers
	route.RegisterRoutes(r, linkController, userController)

	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
