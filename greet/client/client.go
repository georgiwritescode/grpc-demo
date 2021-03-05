package main

import (
	"fmt"
	greetpb "grpc-demo/greet/proto"
	"log"

	"google.golang.org/grpc"
)

func main() {
	log.Println("Hello I'm the client!")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	fmt.Printf("Created client: %f", c)
}
