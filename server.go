package main

import (
	"context"
	"fmt"
	. "go-learning/protobuff"
	"google.golang.org/grpc"
	"net"
)

func main() {
	listener, _ := net.Listen("tcp", ":9090")
	grpcserver := grpc.NewServer()
	RegisterHelloWorldServiceServer(grpcserver, &HelloServer{})
	grpcserver.Serve(listener)
}

type HelloServer struct {
}

func (HelloServer) SayHelloWorld(ctx context.Context, req *HttpRequest) (*HttpResponse, error) {
	fmt.Println("client connect to server")
	fmt.Printf("req.Greeting : %s, req.infos: %s", req.Greeting, req.Infos)
	return &HttpResponse{
		Reply:   "connected successfully",
		Details: []string{"April", "jiangnenghua",},
	}, nil
}
