package grpc

import (
	"context"

	"github.com/Golerplate/contracts/generated/services"
	"github.com/Golerplate/contracts/generated/services/servicesconnect"
	connect_go "github.com/bufbuild/connect-go"
)

type health struct{}

func (h *health) CheckHealth(ctx context.Context, c *connect_go.Request[services.HealthRequest]) (*connect_go.Response[services.HealthResponse], error) {
	return connect_go.NewResponse(&services.HealthResponse{}), nil
}

func NewHealthHandler() servicesconnect.HealthServiceHandler {
	return &health{}
}
