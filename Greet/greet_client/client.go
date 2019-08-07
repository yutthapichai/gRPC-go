package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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
	// doServerStreaming(c)
	doClientStreaming(c)
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

func doServerStreaming(c greetpb.GreetServiceClient) {
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

func doClientStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do client streaming RPC...")
	requests := []*greetpb.LongGreetRequest{
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				Firstname: "Yutdev",
				Lastname:  "Golang",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				Firstname: "Yutdev",
				Lastname:  "Mean stack",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				Firstname: "Yutdev",
				Lastname:  "Mongodb",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				Firstname: "Yutdev",
				Lastname:  "Docker",
			},
		},
	}
	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Printf("Error reading streaming to server %v\n", err)
	}

	for _, req := range requests {
		log.Printf("Starting req ...%v\n", req)
		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Printf("Error receiving streaming %v\n", err)
	}
	fmt.Printf("Result Long Greet Respon %v\n", res)
}
