package main

import (
	"context"
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
	doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	log.Println("Start an unary RPC...")
	greeting := greetpb.Greeting{
		FirstName: "Georgi",
		LastName:  "Ivanov",
	}
	req := &greetpb.GreetRequest{
		Greeting: &greeting,
	}

	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Request while calling Greet. Failed with %v", err)
	}
	log.Printf("Response from Greet: %v", res.Result)
}
