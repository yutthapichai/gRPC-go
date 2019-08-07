package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

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

func (*server) SumMany(req *calpb.SumManyRequest, stream calpb.CalculatorService_SumManyServer) error {
	fmt.Printf("Calculate function was invoked with %v\n", req)
	k := req.GetK()
	N := req.GetN()
	for N > 1 {
		if N%k == 0 {
			res := &calpb.SumManyRespone{
				Result: k,
			}
			N = N / k
			stream.Send(res)
			time.Sleep(1000 * time.Millisecond)
		} else {
			k = k + 1
		}
	}
	return nil
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
