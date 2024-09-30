package config

import (
	"errors"
	"os"
)

const (
	httpHostEnvName        = "HTTP_HOST"
	httpPortEnvName        = "HTTP_PORT"
	httpSwaggerPortEnvName = "HTTP_SWAGGER_PORT"
)

type httpConfig struct {
	Host        string
	Port        uint16
	SwaggerPort uint16
}

// HTTPConfig интерфейс для получения конфигурации HTTP
type HTTPConfig interface {
	GetHTTPHost() string
	GetHTTPPort() uint16
	GetSwaggerPort() uint16
}

// GetHTTPConfig возвращает конфигурацию HTTP
func GetHTTPConfig() (*httpConfig, error) {
	host := os.Getenv(httpHostEnvName)
	if len(host) == 0 {
		return nil, errors.New(httpHostEnvName + " parameter not set")
	}

	port := os.Getenv(httpPortEnvName)
	if len(port) == 0 {
		return nil, errors.New(httpPortEnvName + " parameter not set")
	}

	portUint16, err := stringToUint16(port)
	if err != nil {
		return nil, err
	}

	swaggerPort := os.Getenv(httpSwaggerPortEnvName)
	if len(port) == 0 {
		return nil, errors.New(httpSwaggerPortEnvName + " parameter not set")
	}

	swaggerPortUint16, err := stringToUint16(swaggerPort)
	if err != nil {
		return nil, err
	}

	result := &httpConfig{
		Host:        host,
		Port:        portUint16,
		SwaggerPort: swaggerPortUint16,
	}

	return result, nil
}

func (instance *httpConfig) GetHTTPHost() string {
	return instance.Host
}

func (instance *httpConfig) GetHTTPPort() uint16 {
	return instance.Port
}

func (instance *httpConfig) GetSwaggerPort() uint16 {
	return instance.SwaggerPort
}
