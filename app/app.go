package app

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/fatih/color"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/justbrownbear/microservices_course_chat/controllers/chat_controller"
	grpc_api "github.com/justbrownbear/microservices_course_chat/internal/api/grpc"
	"github.com/justbrownbear/microservices_course_chat/internal/config"
	"github.com/justbrownbear/microservices_course_chat/internal/interceptor"
	pg "github.com/justbrownbear/microservices_course_chat/internal/transaction_manager"
)

var dbPool *pgxpool.Pool
var grpcServer *grpc.Server

var grpcConfig config.GRPCConfig

// InitApp initializes the gRPC server and registers the chat controller.
func InitApp(ctx context.Context, postgresqlConfig config.PostgresqlConfig, grpcConfigInstance config.GRPCConfig) error {
	grpcConfig = grpcConfigInstance
	grpcServer = grpc.NewServer(
		// Прописываем интерцептор валидации для всех запросов
		grpc.UnaryInterceptor(interceptor.ValidateInterceptor),
	)
	reflection.Register(grpcServer)

	var err error

	dbDSN :=
		fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
			postgresqlConfig.GetPostgresHost(),
			postgresqlConfig.GetPostgresPort(),
			postgresqlConfig.GetPostgresDb(),
			postgresqlConfig.GetPostgresUser(),
			postgresqlConfig.GetPostgresPassword())
	dbPool, err = pgxpool.New(ctx, dbDSN)
	if err != nil {
		log.Printf(color.RedString("Unable to connect to database: %v\n"), err)
		return err
	}

	transactionManager := pg.InitTransactionManager(dbPool)
	grpcAPI := grpc_api.InitGrpcAPI(transactionManager)

	chat_controller.InitChatController(grpcServer, grpcAPI)

	return nil
}

// StartApp starts the gRPC server on the provided port.
func StartApp() error {
	grpcProtocol := grpcConfig.GetGrpcProtocol()
	grpcHost := grpcConfig.GetGrpcHost()
	grpcPort := grpcConfig.GetGrpcPort()

	listenAddress := net.JoinHostPort(grpcHost, strconv.Itoa(int(grpcPort)))
	listener, err := net.Listen(grpcProtocol, listenAddress)
	if err != nil {
		log.Printf(color.RedString("Failed to initialize listener: %v"), err)
		return err
	}

	log.Printf(color.GreenString("Starting gRPC server on %s"), listenAddress)

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Printf(color.RedString("Failed to start gRPC server: %v"), err)
		return err
	}

	return nil
}

// StopApp - Остановка приложения
func StopApp() {
	log.Println(color.YellowString("Stopping the application right way..."))

	grpcServer.Stop()
	dbPool.Close()

	log.Println(color.GreenString("Application stopped successfully. Bye."))
}
