package proto

import (
	"context"
)

type Server struct{}

func (s Server) Echo(ctx context.Context, request *Request) (*Response, error) {
	return &Response{Output: request.Input}, nil
}
