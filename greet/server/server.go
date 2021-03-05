package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	greetpb "grpc-demo/greet/proto"
)

type server struct{}

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
