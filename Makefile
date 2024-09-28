include .env.local

LOCAL_BIN:=$(CURDIR)/bin

LOCAL_MIGRATION_DIR=$(MIGRATION_PATH)
LOCAL_MIGRATION_DSN="host=localhost port=$(POSTGRES_PORT) dbname=$(POSTGRES_DB) user=$(POSTGRES_USER) password=$(POSTGRES_PASSWORD) sslmode=disable"

REGISTRY:=cr.selcloud.ru/docker-registry
REGISTRY_USER:=token
REGISTRY_PASS:=CRgAAAAAbxq0rZGxpEw4ny2EuFRK5rGlYUAUqnsZ
PACKAGE_NAME:=chat
PACKAGE_VERSION:=v0.0.2

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.14.0
	GOBIN=$(LOCAL_BIN) go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1.26
	GOBIN=$(LOCAL_BIN) go install golang.org/x/tools/cmd/goimports@latest
	GOBIN=$(LOCAL_BIN) go install github.com/gojuno/minimock/v3/cmd/minimock@latest
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.3
	GOBIN=$(LOCAL_BIN) go install github.com/envoyproxy/protoc-gen-validate@v1.0.4
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.20.0
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.20.0

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go get -u google.golang.org/grpc
	go get -u google.golang.org/grpc/reflection
	go get -u github.com/jackc/pgx/v5
	go get -u github.com/jackc/pgx/v5/pgxpool
	go get -u github.com/gojuno/minimock/v3
	go get -u github.com/brianvoe/gofakeit
	go get -u github.com/brianvoe/gofakeit/v6
	go get -u github.com/stretchr/testify/require
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go get -u github.com/grpc-ecosystem/grpc-gateway/v2/runtime

format:
	find . -name '*.go' -exec $(LOCAL_BIN)/goimports -w {} \;

lint:
	$(LOCAL_BIN)/golangci-lint run ./... --config .golangci.pipeline.yaml

goose-migration-status:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_PATH} postgres "host=$(POSTGRES_HOST) port=$(POSTGRES_PORT) user=$(POSTGRES_USER) password=$(POSTGRES_PASSWORD) dbname=$(POSTGRES_DB) sslmode=disable" status -v

goose-migration-up:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_PATH} postgres "host=$(POSTGRES_HOST) port=$(POSTGRES_PORT) user=$(POSTGRES_USER) password=$(POSTGRES_PASSWORD) dbname=$(POSTGRES_DB) sslmode=disable" up -v

goose-migration-down:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_PATH} postgres "host=$(POSTGRES_HOST) port=$(POSTGRES_PORT) user=$(POSTGRES_USER) password=$(POSTGRES_PASSWORD) dbname=$(POSTGRES_DB) sslmode=disable" down -v

generate:
	make generate-chat-api && \
	make generate-sqlc

generate-chat-api:
	mkdir -p pkg/chat_v1
	protoc \
	--proto_path api/chat_v1 \
	--proto_path vendor.protogen \
	--go_out=pkg/chat_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/chat_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	--validate_out lang=go:pkg/chat_v1 --validate_opt=paths=source_relative \
	--plugin=protoc-gen-validate=bin/protoc-gen-validate \
	--grpc-gateway_out=pkg/chat_v1 --grpc-gateway_opt=paths=source_relative \
	--plugin=protoc-gen-grpc-gateway=bin/protoc-gen-grpc-gateway \
	api/chat_v1/chat.proto

generate-sqlc:
	$(LOCAL_BIN)/sqlc generate

docker-build-and-push:
	docker buildx build --no-cache --platform linux/amd64 -t $(REGISTRY)/$(PACKAGE_NAME):$(PACKAGE_VERSION) .
	docker login -u $(REGISTRY_USER) -p $(REGISTRY_PASS) $(REGISTRY)
	docker push $(REGISTRY)/$(PACKAGE_NAME):$(PACKAGE_VERSION)

dev-env:
	docker run -d --rm --name chat-postgres -e POSTGRES_USER=$(POSTGRES_USER) -e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) -e POSTGRES_DB=$(POSTGRES_DB) -p 5432:5432 -d postgres

.PHONY: test
test:
	go clean -testcache
	go test ./... -covermode count -count 5

vendor-proto:
	@if [ ! -d vendor.protogen/validate ]; then \
		mkdir -p vendor.protogen/validate &&\
		git clone https://github.com/envoyproxy/protoc-gen-validate vendor.protogen/protoc-gen-validate &&\
		mv vendor.protogen/protoc-gen-validate/validate/*.proto vendor.protogen/validate &&\
		rm -rf vendor.protogen/protoc-gen-validate ;\
	fi
	@if [ ! -d vendor.protogen/google ]; then \
		git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
		mkdir -p  vendor.protogen/google/ &&\
		mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
		rm -rf vendor.protogen/googleapis ;\
	fi
	@if [ ! -d vendor.protogen/protoc-gen-openapiv2 ]; then \
		mkdir -p vendor.protogen/protoc-gen-openapiv2/options &&\
		git clone https://github.com/grpc-ecosystem/grpc-gateway vendor.protogen/openapiv2 &&\
		mv vendor.protogen/openapiv2/protoc-gen-openapiv2/options/*.proto vendor.protogen/protoc-gen-openapiv2/options &&\
		rm -rf vendor.protogen/openapiv2 ;\
	fi
