package handler

import "github.com/gofiber/fiber/v2"


func HelloWorldHandler(c *fiber.Ctx) error {
	return c.SendString("Hello World!")
}