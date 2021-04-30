package tests

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func testIndexHello(t *testing.T, app *fiber.App) {
	req, err := http.NewRequest(
		"GET",
		"/api/",
		nil,
	)
	assert.NoError(t, err)
	res, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode)
	// Read the response body
	body, err := ioutil.ReadAll(res.Body)
	assert.NoError(t, err)
	assert.Equal(t, "Hello, World 👋!", string(body))
}
