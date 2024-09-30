package transaction_manager

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate ../../bin/minimock -i TxManager -o ./mocks/ -p transaction_manager_mock -s "_minimock.go"
