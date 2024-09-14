package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/fatih/color"

	"github.com/justbrownbear/microservices_course_chat/app"
	"github.com/justbrownbear/microservices_course_chat/internal/config"
)

// Путь к файлу конфига
var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

func main() {
	ctx := context.Background()

	// Получаем и валидируем конфиг
	grpcConfig, postgresqlConfig, err := getConfig()
	if err != nil {
		log.Printf(color.RedString("Failed to get config: %v"), err)
		os.Exit(1)
	}

	// Инициализируем приложение
	err = app.InitApp(ctx, postgresqlConfig, grpcConfig)
	if err != nil {
		log.Printf(color.RedString("Failed to init app: %v"), err)
		os.Exit(1)
	}
	defer app.StopApp()

	// Создаем канал в котором будем ловить сигналы ОС
	stopChannel := make(chan os.Signal, 1)

	// Регистрируем сигналы, которые будем ловить в канале stopChannel
	signal.Notify(stopChannel, os.Interrupt, syscall.SIGTERM)

	// Чтобы не блокировать основной поток, запускаем сервер в горутине
	go func() {
		err := app.StartApp()
		if err != nil {
			log.Println(color.RedString("Failed to start app: %v"), err)
		}
		// Отправляем SIGTERM в канал, чтобы корректно завершить приложение и закрыть ресурсы
		stopChannel <- syscall.SIGTERM
	}()

	// Как я понимаю, тут мы застрянем до тех пор, пока не придет сигнал
	<-stopChannel

	// А дальше отработает defer
	log.Println("Shutting down app...")
}

func getConfig() (config.GRPCConfig, config.PostgresqlConfig, error) {
	flag.Parse()

	currentDir, err := os.Getwd()
	if err != nil {
		log.Printf(color.RedString("Failed to get current directory: %v"), err)
		return nil, nil, err
	}
	log.Println("Current Directory:", currentDir)

	err = config.Load(configPath)
	if err != nil {
		log.Printf(color.RedString("Failed to load config: %v"), err)
		return nil, nil, err
	}

	grpcConfig, err := config.GetGrpcConfig()
	if err != nil {
		log.Printf(color.RedString("Failed to get gRPC config: %v"), err)
		return nil, nil, err
	}

	postgresqlConfig, err := config.GetPostgresqlConfig()
	if err != nil {
		log.Printf(color.RedString("Failed to get PostgreSQL config: %v"), err)
		return nil, nil, err
	}

	return grpcConfig, postgresqlConfig, nil
}
