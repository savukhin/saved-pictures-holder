package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertToMap(t *testing.T) {
	testCompressedUser := struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
	}{
		ID:       1,
		Username: "someusername",
	}

	testCompressedUserConverted, err := ConvertToMap(testCompressedUser)

	assert.Nil(t, err)
	assert.Equal(t, map[string]interface{}{"id": float64(1), "username": "someusername"}, testCompressedUserConverted)
}

func TestA(t *testing.T) {
	assert.True(t, true)
}
