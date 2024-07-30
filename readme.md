## gRPC Example

Server Address: `grpc://127.0.0.1:50051`

### RPC methods

- `GetOrder`
- `GetOrders`

### Defining the gRPC service

Full Protobuf file in `pkg/proto/service/order/v1/order.proto`

```protobuf
syntax = "proto3";

service OrderService {
  rpc GetOrder (OrderRequest) returns (OrderResponse);
  rpc GetOrders (Empty) returns (OrdersResponse);
}

message OrderRequest {
  int64 order_id = 1;
}

message OrderResponse {
  Order order = 1;
}

message Order {
  int64 order_id = 1;
  string user_name = 2;
  string email = 3;
}

message Empty {}

message OrdersResponse {
  repeated Order orders = 1;
}
```

### GetOrder response

Status code: `0 OK`

```json
{
  "order": {
    "order_id": 3,
    "user_name": "TestUserName3",
    "email": "test3@test.com"
  }
}
```

### GetOrders response

Status code: `0 OK`

```json
{
  "orders": [
    {
      "order_id": 3,
      "user_name": "TestUserName3",
      "email": "test3@test.com"
    },
    {
      "order_id": 2,
      "user_name": "TestUserName2",
      "email": "test2@test.com"
    },
    {
      "order_id": 1,
      "user_name": "TestUserName1",
      "email": "test1@test.com"
    }
  ]
}
```
