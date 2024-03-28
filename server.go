package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/itoldthekettleoff/go-forum/database"
	"github.com/itoldthekettleoff/go-forum/router"
	"github.com/joho/godotenv"
)

var port string = ":3001"

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	} else {
		port = os.Getenv("PORT")
	}
	database.Connect()
}

func main() {
	db, err := database.DBConn.DB()
	if err != nil {
		panic("Error in sql connection.")
	}
	defer db.Close()

	app := fiber.New()
	app.Use(logger.New())
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  200,
			"message": "success",
		})
	})

	router.SetupRoutes(app)
	app.Listen(port)
}
