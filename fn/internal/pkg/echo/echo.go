package echo

import "errors"

var (
	ErrorMissingMessage = errors.New("missing required field Input")
)

type Request struct {
	Input *string `json:"input"`
}

type Response struct {
	Output *string `json:"output"`
}

type Echo struct{}

func (e *Echo) Handle(req Request) (*Response, error) {
	if req.Input == nil || *req.Input == "" {
		return nil, ErrorMissingMessage
	}
	return &Response{Output: req.Input}, nil
}
