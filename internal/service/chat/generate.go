package chat_service

//go:generate cmd /C "mkdir mocks && rmdir /S /Q mocks && mkdir mocks"
//go:generate ../../../bin/minimock -i ChatService -o ./mocks -s "_minimock.go"
