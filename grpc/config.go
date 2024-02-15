package grpc

type GRPCServerConfig struct {
	Port uint16 `env:"GRPC_SERVER_PORT" envDefault:"50051"`
}
