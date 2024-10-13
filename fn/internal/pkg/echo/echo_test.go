package echo_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tmsmr/go-lambda-terraform/fn/internal/pkg/echo"
)

func TestHandle_ReturnsResponseWithOutput_WhenInputIsProvided(t *testing.T) {
	e := &echo.Echo{}
	msg := "a test message"
	req := echo.Request{Input: &msg}

	resp, err := e.Handle(req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, msg, *resp.Output)
}

func TestHandle_ReturnsError_WhenInputIsNil(t *testing.T) {
	e := &echo.Echo{}
	req := echo.Request{Input: nil}

	resp, err := e.Handle(req)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Equal(t, echo.ErrorMissingMessage, err)
}

func TestHandle_ReturnsError_WhenInputIsEmpty(t *testing.T) {
	e := &echo.Echo{}
	msg := ""
	req := echo.Request{Input: &msg}

	resp, err := e.Handle(req)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Equal(t, echo.ErrorMissingMessage, err)
}
