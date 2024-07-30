.DEFAULT_GOAL := gen

gen:
	protoc -I ./pkg/proto \
    --go_out=./pkg/proto_gen/go --go_opt=paths=source_relative \
    --go-grpc_out=./pkg/proto_gen/go --go-grpc_opt=paths=source_relative \
    service/order/v1/order.proto

install:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

run:
	go run cmd/server/main.go
