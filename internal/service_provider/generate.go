package service_provider

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate ../../bin/minimock -i ServiceProvider -o ./mocks/ -p service_provider_mock -s "_minimock.go"
