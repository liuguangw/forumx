package tests

import (
	"github.com/gofiber/fiber/v2"
	"github.com/liuguangw/forumx/cmd"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestAPI(t *testing.T) {
	app := cmd.SetupApp()
	t.Run("index.Hello", func(t *testing.T) {
		testIndexHello(app, t)
	})
}

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
