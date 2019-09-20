package main

import (
	"context"
	"fmt"
	"log"

	pb "greeter/protos"

	"google.golang.org/grpc"
)

func interceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	fmt.Printf("Client sending out request: %v\n", req)
	err := invoker(ctx, method, req, reply, cc, opts...)
	fmt.Printf("Client getting response: %v\n", reply)
	return err
}

func main() {
	conn, err := grpc.Dial(
		"localhost:8080",
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(interceptor))
	if err != nil {
		log.Fatalf("error dialing greeter server: %v", err)
	}
	client := pb.NewGreetServiceClient(conn)

	res, err := client.Greet(context.Background(), &pb.GreetRequest{
		Name: "Brandon",
	})
	if err != nil {
		log.Fatalf("error greeting: %v", err)
	}

	fmt.Println(res.GetMessage())
}
