package main

import (
	"context"
	"fmt"
	"io"
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
	// doUnary(c)
	doStreaming(c)
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

func doStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do Streaming RPC..")
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			Firstname: "Yutdev",
			Lastname:  "Golang",
		},
	}
	resStream, err := c.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error calling greet many time RPC %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Error reading streaming from greet: %v", err)
		}
		log.Printf("Respone from greet many time %v", msg.GetResult())
	}
}
