package main

import (
	"fmt"

	"github.com/devgotech/ticket-booking-v1/config"
	"github.com/devgotech/ticket-booking-v1/db"
	"github.com/devgotech/ticket-booking-v1/handlers"
	"github.com/devgotech/ticket-booking-v1/middlewares"
	"github.com/devgotech/ticket-booking-v1/repositories"
	"github.com/devgotech/ticket-booking-v1/services"
	"github.com/gofiber/fiber/v2"
)

func main() {
	envConfig := config.NewEnvConfig()
	db := db.Init(envConfig, db.DBMigrator)

	app := fiber.New(fiber.Config{
		AppName:      "TicketBooking",
		ServerHeader: "Fiber",
	})

	// Repositories
	eventRepository := repositories.NewEventRepository(db)
	ticketRepository := repositories.NewTicketRepository(db)
	authRepository := repositories.NewAuthRepository(db)

	// service
	authService := services.NewAuthService(authRepository)

	// Routing
	server := app.Group("/api")
	handlers.NewAuthHandler(server.Group("/auth"), authService)

	privateRoutes := server.Use(middlewares.AuthProtected(db))

	// Handlers
	handlers.NewEventHandler(privateRoutes.Group("/event"), eventRepository)
	handlers.NewTicketHandler(privateRoutes.Group("/ticket"), ticketRepository)
	app.Listen(fmt.Sprintf(":" + envConfig.ServerPort))
}
