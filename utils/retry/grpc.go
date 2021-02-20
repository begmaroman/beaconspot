package retry

import (
	"context"
	"time"

	"go.uber.org/zap"

	"google.golang.org/grpc"
)

// UnaryInterceptor is the custom retry unary interceptor
func UnaryInterceptor(logger *zap.Logger, attempts uint, sleep time.Duration) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		return DoFunc(attempts, sleep, func(left uint) error {
			logger.Info("sending RPC request...", zap.Uint("attempts_left", left), zap.String("method", method), zap.String("endpoint", cc.Target()))
			return invoker(ctx, method, req, reply, cc, opts...)
		})
	}
}
