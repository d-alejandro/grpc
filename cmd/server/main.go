package main

import (
	"d-alejandro/grpc/internal/server"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	const port = 50051
	address := fmt.Sprintf(":%d", port)

	listener, listenErr := net.Listen("tcp", address)
	if listenErr != nil {
		log.Fatalf("failed to listen: %v", listenErr)
	}

	grpcServer := grpc.NewServer()
	server.RegisterOrderHandler(grpcServer)

	log.Printf("grpcServer listening at %v", listener.Addr())

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
