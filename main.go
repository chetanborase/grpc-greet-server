package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"

	greet "github.com/chetanborase/grpc-greet-proto/grpc/gen/go/greeting/v1"
)

// GreetService
// Lets implement GreetService by complying service definition given in greet_service.proto file
type GreetService struct {
	greet.UnimplementedGreetServiceServer
}

func (u *GreetService) SayHello(_ context.Context, req *greet.GreetRequest) (*greet.GreetResponse, error) {
	return &greet.GreetResponse{
		Reply: fmt.Sprintf("Hello, %s", req.Name),
	}, nil
}

func main() {
	// create listener that is required for our grpc server
	lis, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	//create new server
	grpcServer := grpc.NewServer()

	//grpc reflection to discover endpoints from our server.
	reflection.Register(grpcServer)

	//we have listener, server and service
	//however the service is still not bounded to server
	//lets do this by Registering out implemented service to the grpc service.
	greet.RegisterGreetServiceServer(grpcServer, &GreetService{})

	log.Println("Starting grpc server.")
	//lets start serving
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to start server with error %v", err)
	}
}
