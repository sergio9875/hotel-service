package main

import (
	"flag"
	"github.com/gofiber/fiber/v2"
	"github.com/sergio9875/hotel-service/api"
)

func main() {
	listenAddr := flag.String("listenAddr", ":4000", "The listen address of the API server")
	flag.Parse()

	app := fiber.New()
	apiV1 := app.Group("/api/v1")

	apiV1.Get("/user", api.HandleUser)
	apiV1.Get("/user/:id", api.HandleGetUserById)
	err := app.Listen(*listenAddr)
	if err != nil {
		return
	}

}
