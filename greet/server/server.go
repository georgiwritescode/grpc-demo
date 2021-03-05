package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	greetpb "grpc-demo/greet/proto"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	log.Default().Printf("Greet func was invoked with %v", req)
	firstName := req.GetGreeting().GetFirstName()

	result := "Hello " + firstName
	res := &greetpb.GreetResponse{
		Result: result,
	}

	return res, nil
}

func main() {
	log.Println("Hello I'm the server!")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
