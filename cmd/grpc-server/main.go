package main

import (
	"log"
	"net"

	helloGrpc "hellogo/internal/grpc"
	"hellogo/proto_gen/helloworld"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	helloworld.RegisterGreeterServer(s, &helloGrpc.Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
