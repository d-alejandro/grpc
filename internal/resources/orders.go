package resources

import orderService "d-alejandro/grpc/pkg/proto_gen/go/service/order/v1"

func GetOrders() map[int64]*orderService.Order {
	return map[int64]*orderService.Order{
		1: {
			OrderId:  1,
			UserName: "TestUserName1",
			Email:    "test1@test.com",
		},
		2: {
			OrderId:  2,
			UserName: "TestUserName2",
			Email:    "test2@test.com",
		},
		3: {
			OrderId:  3,
			UserName: "TestUserName3",
			Email:    "test3@test.com",
		},
	}
}
