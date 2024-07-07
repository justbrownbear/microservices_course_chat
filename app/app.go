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


const GRPC_PROTOCOL = "tcp";


var grpcServer *grpc.Server


func InitApp() {
	grpcServer = grpc.NewServer()
	reflection.Register( grpcServer )

	chat_controller.InitChatController( grpcServer )
}



func StartApp( gRpcPort uint16 ) error {

	listenAddress := ":" + strconv.Itoa( int( gRpcPort ) )
	listener, err := net.Listen( GRPC_PROTOCOL, listenAddress)

	if err != nil {
		log.Printf( color.RedString( "Failed to initialize listener: %v" ), err )
		return err
	}

	log.Printf( color.GreenString( "Starting gRPC server on %s" ), listenAddress )

	err = grpcServer.Serve( listener )

	if err != nil {
		log.Printf( color.RedString( "Failed to start gRPC server: %v" ), err )
		return err
	}

	return nil
}
