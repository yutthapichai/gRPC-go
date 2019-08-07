package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/yutthapichai/gRPC-go/Calculator/calpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello i'm client")
	cc, err := grpc.Dial("localhost:50023", grpc.WithInsecure()) // cc is connect client
	if err != nil {
		log.Fatalf("count not connect %v\n", err)
	}

	defer cc.Close()

	c := calpb.NewCalculatorServiceClient(cc)
	// fmt.Printf("Created client: %f", c)
	// doUnary(c)
	doStreaming(c)
}

func doUnary(c calpb.CalculatorServiceClient) {
	fmt.Println("Starting to do aUnary RPC..")
	req := &calpb.SumRequest{
		FirstNumber:  5,
		SecondNumber: 5,
	}
	res, err := c.Sum(context.Background(), req)

	if err != nil {
		log.Fatalf("Error calling calculator RPC %v", err)
	}
	log.Printf("Respon from Calculator: %v", res)
}

func doStreaming(c calpb.CalculatorServiceClient) {
	fmt.Println("Starting to do Streaming RPC..")
	req := &calpb.SumManyRequest{
		K: 2,
		N: 120,
	}
	resStream, err := c.SumMany(context.Background(), req) // API

	if err != nil {
		log.Fatalf("Error calling Calculate many time RPC %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Error reading streaming from calculate: %v", err)
		}
		log.Printf("Respone from calculate many time %v", msg.GetResult())
	}
}
