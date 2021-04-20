package tests

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func testIndexHello(app *fiber.App, t *testing.T) {
	req, err := http.NewRequest(
		"GET",
		"/api/",
		nil,
	)
	assert.Nil(t, err)
	res, err := app.Test(req, -1)
	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)
	// Read the response body
	body, err := ioutil.ReadAll(res.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello, World 👋!", string(body))
}
