package client

import (
	"context"
	orderService "d-alejandro/grpc/pkg/proto_gen/go/service/order/v1"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"time"
)

type GrpcClient struct {
	orderServiceClient orderService.OrderServiceClient
}

func NewGRPCClient(grpcConnection grpc.ClientConnInterface) *GrpcClient {
	return &GrpcClient{
		orderServiceClient: orderService.NewOrderServiceClient(grpcConnection),
	}
}

func (receiver *GrpcClient) Print() {
	const timeoutDuration = 10 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDuration)
	defer cancel()

	receiver.printOrder(ctx)
	receiver.printOrders(ctx)
}

func (receiver *GrpcClient) printOrder(ctx context.Context) {
	orderRequest := &orderService.OrderRequest{
		OrderId: 3,
	}
	response, err := receiver.orderServiceClient.GetOrder(ctx, orderRequest)

	if err != nil {
		fmt.Printf("could not get order: %v", err)
		return
	}

	data := receiver.convertToString(response)

	fmt.Printf("\nGetOrder: %s", data)
}

func (receiver *GrpcClient) printOrders(ctx context.Context) {
	emptyRequest := &orderService.Empty{}
	response, err := receiver.orderServiceClient.GetOrders(ctx, emptyRequest)

	if err != nil {
		fmt.Printf("could not get orders: %v", err)
		return
	}

	data := receiver.convertToString(response)

	fmt.Printf("\n\nGetOrders: %s", data)
}

func (receiver *GrpcClient) convertToString(data any) string {
	bytes, _ := json.MarshalIndent(data, "", "  ")
	return string(bytes)
}
