package main

import (
	"d-alejandro/grpc/internal/client"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	const port = 50051
	address := fmt.Sprintf(":%d", port)

	connection, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer func() {
		if err := connection.Close(); err != nil {
			log.Fatalf("could not close connection: %v", err)
		}
	}()

	grpcClient := client.NewGRPCClient(connection)
	grpcClient.Print()
}
