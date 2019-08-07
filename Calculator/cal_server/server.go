package main

import (
	"context"
	"fmt"
	"io"
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

func (*server) SumLong(stream calpb.CalculatorService_SumLongServer) error {
	fmt.Printf("Starting to do Streaming RPC..\n")
	sum := int32(0)
	count := 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			average := float64(sum) / float64(count)
			return stream.SendAndClose(&calpb.SumLongRespone{
				LongResult: average,
			})
		}
		if err != nil {
			log.Printf("Error reading streaming from greet: %v", err)
		}
		log.Printf("Respone from amount Long time %v\n", req.GetN())
		sum += req.GetN()
		count++
	}

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
