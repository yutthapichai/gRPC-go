package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func main() {
	fmt.Println("Calculator server")

	lis, err := net.Listen("tcp", "0.0.0.0:50023")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	// calculatorpb..RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}
