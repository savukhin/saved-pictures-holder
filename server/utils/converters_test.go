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

	var testCompressedUserConverted = ConvertToMap(testCompressedUser)

	assert.Equal(t, map[string]interface{}{"ID": 1, "Username": "someusername"}, testCompressedUserConverted)
}

func TestA(t *testing.T) {
	assert.True(t, true)
}
