package main

import (
	"context"
	"fmt"
	pb "greeter/protos"
	"log"
	"net"

	"google.golang.org/grpc"
)

type GreetServer struct{}

func interceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	fmt.Printf("Server received request: %v\n", req)
	res, err := handler(ctx, req)
	fmt.Printf("Server sent response: %v\n", res)
	return res, err
}

// Greet ...
func (g *GreetServer) Greet(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	return &pb.GreetResponse{
		Message: fmt.Sprintf("Hello, %s!", req.GetName()),
	}, nil
}

func main() {
	server := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor),
	)
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("couldn't create listener: %v", err)
	}
	pb.RegisterGreetServiceServer(server, &GreetServer{})
	if err := server.Serve(lis); err != nil {
		log.Fatalf("error serving greet requests: %v", err)
	}
}
