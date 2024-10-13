package main

import (
	"context"
	"errors"
	"github.com/tmsmr/go-lambda-terraform/fn/internal/pkg/echo"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

var (
	ErrorMissingFoo = errors.New("missing required environment variable FOO")
	e               = &echo.Echo{}
)

func handleEcho(_ context.Context, req echo.Request) (*echo.Response, error) {
	if os.Getenv("FOO") == "" {
		return nil, ErrorMissingFoo
	}

	return e.Handle(req)
}

func main() {
	lambda.Start(handleEcho)
}
