package interceptor

import (
	"context"

	"google.golang.org/grpc"
)


type validator interface {
	Validate() error
}


// ValidateRequest валидирует запрос
func ValidateInterceptor(
	ctx context.Context,
	request interface{},
	info *grpc.UnaryServerInfo,
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
