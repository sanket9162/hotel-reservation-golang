package main

import (
	"flag"

	"github.com/gofiber/fiber/v2"
	"github.com/sanket9162/hotel-reservation/api"
)

func main() {
	listenAddr := flag.String("listenAddr", "localhost:5000", "The listen address of the API server")
	flag.Parse()

	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("sanket")
	})

	apiv1.Get("/user", api.HandleGetUsers)
	apiv1.Get("/user/:id", api.HandleGetUser)

	app.Listen(*listenAddr)
}
