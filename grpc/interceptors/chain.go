package grpc

import (
	"time"

	"github.com/Golerplate/pkg/grpc/interceptors/logger"
	"github.com/Golerplate/pkg/grpc/interceptors/recover"
	"github.com/Golerplate/pkg/grpc/interceptors/timeout"
	"github.com/bufbuild/connect-go"
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
