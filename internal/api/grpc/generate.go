package grpc_api

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate ../../../bin/minimock -i GrpcAPI -o ./mocks/ -s "_minimock.go"
