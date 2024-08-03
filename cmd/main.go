package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/fatih/color"

	"github.com/justbrownbear/microservices_course_chat/app"
	"github.com/justbrownbear/microservices_course_chat/internal/config"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

func main() {
	ctx := context.Background()

	// Получаем и валидируем конфиг
	grpcConfig, postgresqlConfig := getConfig()

	// Инициализируем
	app.InitApp(ctx, postgresqlConfig)
	defer app.StopApp()

	gRPCProtocol := grpcConfig.GetGrpcProtocol()
	gRPCHost := grpcConfig.GetGrpcHost()
	gRPCPort := grpcConfig.GetGrpcPort()
	err := app.StartApp(gRPCProtocol, gRPCHost, gRPCPort)
	if err != nil {
		log.Fatalf(color.RedString("Failed to start app: %v", err))
	}
}

func getConfig() (config.GRPCConfig, config.PostgresqlConfig) {
	flag.Parse()

	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf(color.RedString("Failed to get current directory: %v", err))
	}
	log.Println("Current Directory:", currentDir)

	err = config.Load(configPath)
	if err != nil {
		log.Fatalf(color.RedString("Failed to load config: %v", err))
	}

	grpcConfig, err := config.GetGrpcConfig()
	if err != nil {
		log.Fatalf(color.RedString("Failed to get gRPC config: %v", err))
	}

	postgresqlConfig, err := config.GetPostgresqlConfig()
	if err != nil {
		log.Fatalf(color.RedString("Failed to get PostgreSQL config: %v", err))
	}

	return grpcConfig, postgresqlConfig
}
