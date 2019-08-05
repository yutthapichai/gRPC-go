package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/yutthapichai/gRPC-go/Calculator/calpb"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calpb.SumRequest) (*calpb.SumRespone, error) {
	fmt.Printf("Calculator function was API with %v", req)
	firstNumber := req.GetFirstNumber()
	secondNumber := req.GetSecondNumber()
	res := &calpb.SumRespone{
		SumResult: firstNumber + secondNumber,
	}
	return res, nil
}

func main() {
	fmt.Println("Calculator server")

	lis, err := net.Listen("tcp", "0.0.0.0:50023")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}
