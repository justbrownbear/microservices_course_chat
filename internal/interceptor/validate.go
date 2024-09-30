package interceptor

import (
	"context"

	"google.golang.org/grpc"
)

type validator interface {
	Validate() error
}

// ValidateInterceptor валидирует запрос
func ValidateInterceptor(
	ctx context.Context,
	request interface{},
	_ *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	val, ok := request.(validator)
	if ok {
		err := val.Validate()
		if err != nil {
			return nil, err
		}
	}

	return handler(ctx, request)
}
