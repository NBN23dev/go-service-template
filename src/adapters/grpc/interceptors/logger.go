package interceptors

import (
	"context"
	"encoding/json"

	"github.com/NBN23dev/go-service-template/src/plugins/logger"
	"google.golang.org/grpc"
)

// InterceptUnary
func InterceptUnary(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	res, err := handler(ctx, req)

	if err != nil {
		body, _ := json.Marshal(req)

		payload := map[string]string{
			"name": info.FullMethod,
			"req":  string(body),
		}

		logger.Error(err.Error(), payload)
	}

	return res, err
}

func InterceptStream(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	err := handler(srv, ss)

	if err != nil {
		logger.Error(err.Error(), nil)
	}

	return err
}
