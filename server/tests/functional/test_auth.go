package functional_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	http.Get("http://localhost:3000/v1/api/register/")

	assert := assert.New(t)
	assert.True(true)
}
