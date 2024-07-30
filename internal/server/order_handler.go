package server

import (
	"cmp"
	"context"
	"d-alejandro/grpc/internal/resources"
	orderService "d-alejandro/grpc/pkg/proto_gen/go/service/order/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"slices"
)

type orderHandler struct {
	orderService.UnimplementedOrderServiceServer
	orders map[int64]*orderService.Order
}

func RegisterOrderHandler(grpcServer *grpc.Server) {
	orderService.RegisterOrderServiceServer(grpcServer, newOrderHandler())
}

func newOrderHandler() *orderHandler {
	orderHandler := &orderHandler{}
	orderHandler.orders = resources.GetOrders()
	return orderHandler
}

func (receiver *orderHandler) GetOrder(
	context context.Context,
	request *orderService.OrderRequest,
) (*orderService.OrderResponse, error) {
	orderId := request.GetOrderId()

	const emptyValue = 0

	if orderId == emptyValue {
		return nil, status.Errorf(codes.InvalidArgument, "order_id is required")
	}

	order, found := receiver.orders[orderId]
	if !found {
		return nil, status.Errorf(codes.NotFound, "order not found")
	}

	orderResponse := &orderService.OrderResponse{
		Order: order,
	}

	return orderResponse, nil
}

func (receiver *orderHandler) GetOrders(
	context.Context,
	*orderService.Empty,
) (*orderService.OrdersResponse, error) {
	orders := make([]*orderService.Order, 0, len(receiver.orders))

	for _, order := range receiver.orders {
		orders = append(orders, order)
	}

	slices.SortStableFunc(orders, func(a, b *orderService.Order) int {
		return -1 * cmp.Compare(a.GetOrderId(), b.GetOrderId())
	})

	ordersResponse := &orderService.OrdersResponse{
		Orders: orders,
	}
	return ordersResponse, nil
}
