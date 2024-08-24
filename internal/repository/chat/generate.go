package chat_repository

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate ../../../bin/minimock -i ChatRepository -o ./mocks -p chat_repository_mock -s "_minimock.go"
