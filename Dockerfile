FROM golang:1.21.12-bookworm AS builder

COPY . /source/
WORKDIR /source/

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/chat ./cmd/main.go

FROM alpine:3.20.2

WORKDIR /root/
COPY --from=builder /source/bin/chat .

CMD ["./chat"]
