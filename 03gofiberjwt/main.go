package main

import (
	"os"

	"gofiberjwt/routes"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/joho/godotenv"
)

var (
	app = fiber.New()
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
	app.Use(requestid.New())
	app.Use(requestid.New(requestid.Config{
		Header: "Test-Service-Header",
		Generator: func() string {
			return utils.UUID()
		},
	}))
	app.Use(logger.New(logger.Config{
		Format:     "${pid} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "Asia/Bangkok",
	}))
}

func main() {
	//not AuthorizationRequired
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Post("/login", routes.Auth)

	//AuthorizationRequired Action
	app.Use(routes.AuthorizationRequired())

	//need AuthorizationRequired
	app.Get("/profile", routes.Profile)
	//end AuthorizationRequired

	err := app.Listen(":" + os.Getenv("APP_PORT"))
	if err != nil {
		panic(err)
	}
}
