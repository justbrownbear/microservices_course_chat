package app

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	"github.com/fatih/color"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/justbrownbear/microservices_course_chat/controllers/chat_controller"
	grpc_api "github.com/justbrownbear/microservices_course_chat/internal/api/grpc"
	"github.com/justbrownbear/microservices_course_chat/internal/config"
	"github.com/justbrownbear/microservices_course_chat/internal/interceptor"
	pg "github.com/justbrownbear/microservices_course_chat/internal/transaction_manager"
	"github.com/justbrownbear/microservices_course_chat/pkg/chat_v1"
)

var dbPool *pgxpool.Pool
var grpcConfig config.GRPCConfig
var grpcServer *grpc.Server
var httpServer *http.Server

// InitApp initializes the gRPC server and registers the chat controller.
func InitApp(
	ctx context.Context,
	postgresqlConfig config.PostgresqlConfig,
	grpcConfigInstance config.GRPCConfig,
	httpConfigInstance config.HTTPConfig,
) error {
	var err error

	grpcConfig = grpcConfigInstance

	grpcServer = initGrpcServer()
	httpServer, err = initHTTPServer(ctx, httpConfigInstance)
	if err != nil {
		return err
	}

	dbPool, err = initPostgreSQLPool(ctx, postgresqlConfig)
	if err != nil {
		return err
	}

	transactionManager := pg.InitTransactionManager(dbPool)
	grpcAPI := grpc_api.InitGrpcAPI(transactionManager)

	chat_controller.InitChatController(grpcServer, grpcAPI)

	return nil
}

// StartApp starts the gRPC server on the provided port.
func StartApp() error {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(2)

	go startGrpcServer(&waitGroup)
	go startHTTPServer(&waitGroup)

	waitGroup.Wait()

	return nil
}

// StopApp - Остановка приложения
func StopApp() {
	log.Println(color.YellowString("Stopping the application right way..."))

	grpcServer.Stop()
	dbPool.Close()

	log.Println(color.GreenString("Application stopped successfully. Bye."))
}

func initGrpcServer() *grpc.Server {
	grpcServerInstance := grpc.NewServer(
		// Прописываем интерцептор валидации для всех запросов
		grpc.UnaryInterceptor(interceptor.ValidateInterceptor),
	)
	reflection.Register(grpcServerInstance)

	return grpcServerInstance
}

func startGrpcServer(waitGroup *sync.WaitGroup) {
	waitGroup.Add(1)
	defer waitGroup.Done()

	grpcProtocol := grpcConfig.GetGrpcProtocol()
	listenAddress := getGrpcAddress()
	listener, err := net.Listen(grpcProtocol, listenAddress)
	if err != nil {
		log.Printf(color.RedString("Failed to initialize listener: %v"), err)
		return
	}

	log.Printf(color.GreenString("Starting gRPC server on %s"), listenAddress)

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Printf(color.RedString("Failed to start gRPC server: %v"), err)
		return
	}
}

func initPostgreSQLPool(
	ctx context.Context,
	postgresqlConfig config.PostgresqlConfig,
) (*pgxpool.Pool, error) {
	dbDSN :=
		fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
			postgresqlConfig.GetPostgresHost(),
			postgresqlConfig.GetPostgresPort(),
			postgresqlConfig.GetPostgresDb(),
			postgresqlConfig.GetPostgresUser(),
			postgresqlConfig.GetPostgresPassword())
	dbPool, err := pgxpool.New(ctx, dbDSN)
	if err != nil {
		log.Printf(color.RedString("Unable to connect to database: %v\n"), err)
		return nil, err
	}

	return dbPool, nil
}

func initHTTPServer(
	ctx context.Context,
	httpConfigInstance config.HTTPConfig,
) (*http.Server, error) {
	multiplexer := runtime.NewServeMux()

	options := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	grpcAddress := getGrpcAddress()

	err := chat_v1.RegisterChatV1HandlerFromEndpoint(ctx, multiplexer, grpcAddress, options)
	if err != nil {
		return nil, err
	}

	httpHost := httpConfigInstance.GetHTTPHost()
	httpPort := httpConfigInstance.GetHTTPPort()
	listenAddress := net.JoinHostPort(httpHost, strconv.Itoa(int(httpPort)))

	httpServerInstance := &http.Server{
		Addr:              listenAddress,
		Handler:           multiplexer,
		ReadHeaderTimeout: 5 * time.Second,
	}

	return httpServerInstance, nil
}

func startHTTPServer(waitGroup *sync.WaitGroup) {
	waitGroup.Add(1)
	defer waitGroup.Done()

	log.Printf(color.GreenString("Starting HTTP server on %s"), httpServer.Addr)

	err := httpServer.ListenAndServe()
	if err != nil {
		log.Printf(color.RedString("Failed to start HTTP server: %v"), err)
		return
	}
}

func getGrpcAddress() string {
	grpcHost := grpcConfig.GetGrpcHost()
	grpcPort := grpcConfig.GetGrpcPort()

	listenAddress := net.JoinHostPort(grpcHost, strconv.Itoa(int(grpcPort)))

	return listenAddress
}
