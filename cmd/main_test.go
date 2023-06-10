package main

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/Sermanes/learning.fiber-api-example/handler"
	"github.com/stretchr/testify/assert"

	"github.com/gofiber/fiber/v2"
)

func TestMain(t *testing.T) {
	app := fiber.New()
	t.Run("Test Hello World", func(t *testing.T) {
		// Set expected value
		want := "Hello World!"

		// Set the endpoint with handler
		app.Get("/", handler.HelloWorldHandler)

		// Make the request
		request := httptest.NewRequest("GET", "http://test.com/", nil)
		response, _ := app.Test(request)

		// Compare the results
		assert.Equal(t, fiber.StatusOK, response.StatusCode)
		body, _ := ioutil.ReadAll(response.Body)
		assert.Equal(t, want, string(body))
	})
}