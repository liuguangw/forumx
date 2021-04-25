package tools

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMd5String(t *testing.T) {
	hashValue := Md5String("123456")
	assert.Equal(t, "e10adc3949ba59abbe56e057f20f883e", hashValue)
}
