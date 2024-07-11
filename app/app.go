package app

import (
	"log"
	"net"
	"strconv"

	"github.com/fatih/color"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/justbrownbear/microservices_course_chat/controllers/chat_controller"
)

var grpcServer *grpc.Server

// InitApp initializes the gRPC server and registers the chat controller.
func InitApp() {
	grpcServer = grpc.NewServer()
	reflection.Register(grpcServer)

	chat_controller.InitChatController(grpcServer)
}

// StartApp starts the gRPC server on the provided port.
func StartApp(grpcProtocol string, grpcPort uint16) error {
	listenAddress := ":" + strconv.Itoa(int(grpcPort))
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
