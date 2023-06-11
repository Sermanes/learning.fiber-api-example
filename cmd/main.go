package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Sermanes/learning.fiber-api-example/handler"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config := fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		ServerHeader:  "HTTP Server",
		AppName:       "Learning Project",
		ReadTimeout:   20 * time.Second,
	}

	app := fiber.New(config)

	// Allow only one connection per IP
	app.Server().MaxConnsPerIP = 2

	// This always happens
	app.Use(func(c *fiber.Ctx) error {
		log.Println("Hi, I always show up!")
		return c.Next()
	})

	// This always happens when match requests starting with /api or /static
	app.Use([]string{"/api", "/static"}, func(c *fiber.Ctx) error {
		log.Println("I am an api or static endpoint, right?")
		return c.Next()
	})

	// Multiple handlers in one Use
	app.Use("/mandatory", func(c *fiber.Ctx) error {
		c.Set("X-Custom-Header", "test")
		return c.Next()
	}, func(c *fiber.Ctx) error {
		return c.Next()
	})

	app.Get("/", handler.HelloWorldHandler)

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
		Compress:      true,
		ByteRange:     true,
		CacheDuration: 10 * time.Second,
		MaxAge:        20,
	})

	app.Get("/sleep", func(c *fiber.Ctx) error {
		time.Sleep(10 * time.Second)
		return c.SendString("Sleeping")
	}).Name("Sleep")

	app.Get("/shutdown", func(c *fiber.Ctx) error {
		return app.Shutdown()
	}).Name("Shutdown")

	// Enpoint's map
	app.Get("/map", func(c *fiber.Ctx) error {
		return c.JSON(c.App().Stack())
	})

	// Get url params
	app.Get("/:name/:surname", func(c *fiber.Ctx) error {
		params := c.AllParams()
		return c.SendString(fmt.Sprintf("Hello %s %s", params["name"], params["surname"]))
	})

	app.Get("/cookie", func(c *fiber.Ctx) error {
		c.Cookie(&fiber.Cookie{
			Name: "token",
			Value: "1234",
			Expires: time.Now().Add(24 * time.Hour),
		})

		return c.SendString("We've set a cookie")
	})

	app.Get("/clear-cookies", func(c *fiber.Ctx) error {
		c.ClearCookie("token")

		return c.SendString("We've cleaned the cookie")
	})

	app.Listen(":3000")
}
