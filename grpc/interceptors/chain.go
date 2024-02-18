package grpc

import (
	"time"

	"github.com/bufbuild/connect-go"
	"github.com/golerplate/pkg/grpc/interceptors/logger"
	"github.com/golerplate/pkg/grpc/interceptors/recover"
	"github.com/golerplate/pkg/grpc/interceptors/timeout"
)

type InterceptorsChain []connect.Interceptor

func ServerDefaultChain() InterceptorsChain {
	return []connect.Interceptor{
		recover.RecoverInterceptor(),
	}
}

func ClientDefaultChain() InterceptorsChain {
	return []connect.Interceptor{
		timeout.TimeoutInterceptor(5 * time.Second),
		logger.ClientLoggerInterceptor(),
	}
}

func ClientConfigurableChain(t time.Duration) InterceptorsChain {
	return []connect.Interceptor{
		timeout.TimeoutInterceptor(t),
		logger.ClientLoggerInterceptor(),
	}
}
