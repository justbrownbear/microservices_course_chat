package chat_service

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate ../../../bin/minimock -i ChatService -o ./mocks -p chat_service_mock -s "_minimock.go"
