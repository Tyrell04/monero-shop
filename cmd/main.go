package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"log"
	_ "monero-shop-api/docs"
	"monero-shop-api/internal/adapter/postgres"
	"monero-shop-api/internal/core/util"
	"monero-shop-api/internal/exception"
	"os"
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8000
// @BasePath /
func main() {
	// Init config
	config := util.New().Get()
	// Init logger
	logFile, err := os.OpenFile("log", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)

	exception.PanicLogging(err)

	defer logFile.Close()

	// Set log out put and enjoy :)
	log.SetOutput(logFile)

	// optional: log date-time, filename, and line number
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	// Init database
	ctx := context.Background()
	db, err := postgres.New(ctx, &config)
	exception.PanicLogging(err)

	defer db.Close()

	app := fiber.New()

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	// add sagger docs for health endpoint

	app.Get("/health", healthHandler)

	app.Listen(":8000")
}

// healthHandler godoc
// @Summary Show the status of the service
// @Description get the status of the service
// @Tags health
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"OK"
// @Router /health [get]
func healthHandler(c *fiber.Ctx) error {
	return c.SendString("OK")
}
