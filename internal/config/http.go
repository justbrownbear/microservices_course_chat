package config

import (
	"errors"
	"os"
)

const (
	httpHostEnvName     = "HTTP_HOST"
	httpPortEnvName     = "HTTP_PORT"
)

type httpConfig struct {
	Host     string
	Port     uint16
}

// HttpConfig интерфейс для получения конфигурации HTTP
type HttpConfig interface {
	GetHttpHost() string
	GetHttpPort() uint16
}

// GetHttpConfig возвращает конфигурацию HTTP
func GetHttpConfig() (*httpConfig, error) {
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

	result := &httpConfig{
		Host:     host,
		Port:     portUint16,
	}

	return result, nil
}

func (instance *httpConfig) GetHttpHost() string {
	return instance.Host
}

func (instance *httpConfig) GetHttpPort() uint16 {
	return instance.Port
}
