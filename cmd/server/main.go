package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/qahta0/media-uploader-microservice/api"
)

func main() {
	engine := html.New("../../views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	api.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
