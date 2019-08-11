package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/yutthapichai/gRPC-go/Greet/greetpb"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet function was invoked with %v\n", req)
	firstname := req.GetGreeting().GetFirstname()
	result := "Hello " + firstname
	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res, nil
}

func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	fmt.Printf("GreetManyTimes function was invoked with %v\n", req)
	firstname := req.GetGreeting().GetFirstname()
	for i := 1; i < 10; i++ {
		result := "Hello " + firstname + " number " + strconv.Itoa(i)
		res := &greetpb.GreetManyTimesResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}

func (*server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	fmt.Printf("Starting to do Streaming RPC..\n")
	result := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&greetpb.LongGreetResponse{
				Result: result,
			})
		}
		if err != nil {
			log.Printf("Error reading streaming from greet: %v", err)
		}
		log.Printf("Respone from greet many time %v\n", req.GetGreeting())
		firstname := req.GetGreeting().GetFirstname()
		lastname := req.GetGreeting().GetLastname()
		result += "Hello " + firstname + " " + lastname + "!\n"
	}

}

func (*server) GreetEveryTime(stream greetpb.GreetService_GreetEveryTimeServer) error {
	fmt.Printf("starting to client streaming and server streaming\n")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Printf("Error while reading client: %v", err)
			return err
		}
		firtsname := req.GetGreeting().GetFirstname()
		lastname := req.GetGreeting().GetLastname()
		result := "Hello " + firtsname + " " + lastname
		errError := stream.Send(&greetpb.GreetEveryTimeResponse{
			Result: result,
		})
		if errError != nil {
			log.Printf("Error while sending data to client %v", err)
			return err
		}
	}
}

func main() {
	fmt.Println("Hello world server")

	lis, err := net.Listen("tcp", "0.0.0.0:50023")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}
