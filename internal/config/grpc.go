package config

import (
	"errors"
	"os"
)

const (
	grpcProtocolEnvName = "GRPC_PROTOCOL"
	grpcHostEnvName     = "GRPC_HOST"
	grpcPortEnvName     = "GRPC_PORT"

	grpcDefaultProtocol = "tcp"
)

type grpcConfig struct {
	Protocol string
	Host     string
	Port     uint16
}

// GRPCConfig интерфейс для получения конфигурации gRPC
type GRPCConfig interface {
	GetGrpcProtocol() string
	GetGrpcHost() string
	GetGrpcPort() uint16
}

// GetGrpcConfig возвращает конфигурацию gRPC
func GetGrpcConfig() (GRPCConfig, error) {
	protocol := os.Getenv(grpcProtocolEnvName)
	if len(protocol) == 0 {
		protocol = grpcDefaultProtocol
	}

	host := os.Getenv(grpcHostEnvName)

	port := os.Getenv(grpcPortEnvName)
	if len(port) == 0 {
		return nil, errors.New(grpcPortEnvName + " parameter not set")
	}

	portUint16, err := stringToUint16(port)
	if err != nil {
		return nil, err
	}

	result := &grpcConfig{
		Protocol: protocol,
		Host:     host,
		Port:     portUint16,
	}

	return result, nil
}

func (c *grpcConfig) GetGrpcProtocol() string {
	return c.Protocol
}

func (c *grpcConfig) GetGrpcHost() string {
	return c.Host
}

func (c *grpcConfig) GetGrpcPort() uint16 {
	return c.Port
}
