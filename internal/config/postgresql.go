package config

import (
	"errors"
	"os"
)

const (
	postgresUserEnvName = "POSTGRES_USER"
	// nolint:gosec
	postgresPasswordEnvName = "POSTGRES_PASSWORD"
	postgresHostEnvName     = "POSTGRES_HOST"
	postgresPortEnvName     = "POSTGRES_PORT"
	postgresDbEnvName       = "POSTGRES_DB"
)

type postgresConfig struct {
	User     string
	Password string
	Host     string
	Port     uint16
	Db       string
}

// PostgresqlConfig интерфейс для получения конфигурации PostgreSQL
type PostgresqlConfig interface {
	GetPostgresUser() string
	GetPostgresPassword() string
	GetPostgresHost() string
	GetPostgresPort() uint16
	GetPostgresDb() string
}

// GetPostgresqlConfig возвращает конфигурацию PostgreSQL
func GetPostgresqlConfig() (PostgresqlConfig, error) {
	user := os.Getenv(postgresUserEnvName)
	if len(user) == 0 {
		return nil, errors.New(postgresUserEnvName + " parameter not set")
	}

	password := os.Getenv(postgresPasswordEnvName)
	if len(password) == 0 {
		return nil, errors.New(postgresPasswordEnvName + " parameter not set")
	}

	host := os.Getenv(postgresHostEnvName)
	if len(host) == 0 {
		return nil, errors.New(postgresHostEnvName + " parameter not set")
	}

	port := os.Getenv(postgresPortEnvName)
	if len(port) == 0 {
		return nil, errors.New(postgresPortEnvName + " parameter not set")
	}

	portUint16, err := stringToUint16(port)
	if err != nil {
		return nil, err
	}

	db := os.Getenv(postgresDbEnvName)
	if len(db) == 0 {
		return nil, errors.New(postgresDbEnvName + " parameter not set")
	}

	result := &postgresConfig{
		User:     user,
		Password: password,
		Host:     host,
		Port:     portUint16,
		Db:       db,
	}

	return result, nil
}

func (c *postgresConfig) GetPostgresUser() string {
	return c.User
}

func (c *postgresConfig) GetPostgresPassword() string {
	return c.Password
}

func (c *postgresConfig) GetPostgresHost() string {
	return c.Host
}

func (c *postgresConfig) GetPostgresPort() uint16 {
	return c.Port
}

func (c *postgresConfig) GetPostgresDb() string {
	return c.Db
}
