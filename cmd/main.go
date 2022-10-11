package main

import (
	"log"

	_ "github.com/adilsonmenechini/goapi/docs"
	"github.com/adilsonmenechini/goapi/internal/routes"
	"github.com/gofiber/fiber/v2"
	swagger "github.com/gofiber/swagger"
)

func main() {
	internal := fiber.New()
	// internal.Use(cors.New())
	internal.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Welcome to the clean-architecture !"))
	})
	api := internal.Group("/api")
	api.Get("/docs/*", swagger.HandlerDefault)
	api.Get("/docs/*", swagger.New(swagger.Config{ // custom
		URL:         "http://example.com/doc.json",
		DeepLinking: false,
	}))
	routes.UserRouter(api)

	log.Fatal(internal.Listen(":3000"))
}
