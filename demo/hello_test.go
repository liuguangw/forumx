package demo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHello(t *testing.T) {
	assertTool := assert.New(t)
	assertTool.Equal(123,123, "they should be equal")
	assertTool.Equal(4, hello(2))
}
