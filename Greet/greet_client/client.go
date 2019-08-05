package main

import (
	"context"
	"fmt"
	"log"

	"github.com/yutthapichai/gRPC-go/Greet/greetpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello i'm client")
	cc, err := grpc.Dial("localhost:50023", grpc.WithInsecure()) // cc is connect client
	if err != nil {
		log.Fatalf("count not connect %v", err)
	}

	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	// fmt.Printf("Created client: %f", c)
	doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do aUnary RPC..")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			Firstname: "Yut",
			Lastname:  "Dev",
		},
	}
	res, err := c.Greet(context.Background(), req)

	if err != nil {
		log.Fatalf("Error calling greet RPC %v", err)
	}
	log.Printf("Respon from greet: %v", res)
}
