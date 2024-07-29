.DEFAULT_GOAL := gen

gen:
	protoc -I ./pkg/proto \
    --go_out=./pkg/proto_gen/go --go_opt=paths=source_relative \
    service/order/v1/order.proto
