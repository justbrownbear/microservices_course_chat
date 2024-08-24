package service_provider

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate ../../bin/minimock -i ServiceProvider -o ./mocks/ -s "_minimock.go"
