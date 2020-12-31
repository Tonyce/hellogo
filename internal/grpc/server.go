package grpc

import (
	"context"
	"log"

	"hellogo/proto_gen/helloworld"
)

const (
	port = ":50051"
)

// Server is used to implement helloworld.GreeterServer.
type Server struct {
	helloworld.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *Server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &helloworld.HelloReply{Message: "Hello " + in.GetName()}, nil
}
