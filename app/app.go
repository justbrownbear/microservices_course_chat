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
	"github.com/justbrownbear/microservices_course_chat/controllers/user_controller"
	grpc_api "github.com/justbrownbear/microservices_course_chat/internal/api/grpc"
	"github.com/justbrownbear/microservices_course_chat/internal/config"
)

var dbPool *pgxpool.Pool
var grpcServer *grpc.Server

// InitApp initializes the gRPC server and registers the chat controller.
func InitApp(ctx context.Context, postgresqlConfig config.PostgresqlConfig) {
	grpcServer = grpc.NewServer()
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
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	grpcAPI := grpc_api.InitGrpcAPI(dbPool)

	chat_controller.InitChatController(grpcServer, grpcAPI)
	user_controller.InitUserController(grpcServer, grpcAPI)
}

// StartApp starts the gRPC server on the provided port.
func StartApp(
	grpcProtocol string,
	grpcHost string,
	grpcPort uint16,
) error {
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
