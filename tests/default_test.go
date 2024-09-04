package tests

import (
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestHomeRoute(t *testing.T) {
	app := fiber.New()

	req, _ := http.NewRequest("GET", "/", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, 200, resp.StatusCode)
}
