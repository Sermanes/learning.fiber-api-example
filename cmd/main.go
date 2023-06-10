package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config := fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		ServerHeader:  "HTTP Server",
		AppName:       "Learning Project",
	}

	app := fiber.New(config)

	app.Use(func (c *fiber.Ctx) error {
		log.Println("Hi, I always show up!")
		return c.Next()
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	// Using mandatory params
	app.Get("/mandatory/:name", func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprintf("Hello %s!", c.Params("name")))
	})

	// Using optional parameter
	app.Get("/optional/:name?", func(c *fiber.Ctx) error {
		if c.Params("name") == "" {
			return c.SendString("What's your name?")
		}

		return c.SendString(fmt.Sprintf("Hello %s", c.Params("name")))
	})

	// Wildcards
	app.Get("/api/*", func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprintf("API Path is: %s", c.Params("*")))
	})

	// Static file
	app.Static("/static", "./public/static")

	// Return an error response
	app.Get("/error", func(c *fiber.Ctx) error {
		return fiber.NewError(500, "Hello Internal Server Error")
	})

	app.Static("/image", "./files/fondo.jpg", fiber.Static{
		Compress: true,
		ByteRange: true,
		CacheDuration: 10 * time.Second,
		MaxAge: 20,
	})

	app.Listen(":3000")
}
