FROM golang:latest AS builder

COPY . /source/
WORKDIR /source/

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/chat ./cmd/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /source/bin/chat .

CMD ["./chat"]
