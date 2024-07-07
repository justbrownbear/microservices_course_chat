package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/justbrownbear/microservices_course_chat/app"
)


const GRPC_PORT = 9098;


func main() {
	app.InitApp()
	err := app.StartApp( GRPC_PORT );

	if err != nil {
		fmt.Println( color.RedString( "Failed to start app: %v", err ) )
	}
}
