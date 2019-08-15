package main

import (
	"context"
	pb "go-learning/protobuff"
	"google.golang.org/grpc"
)

func main() {
	conn, _ := grpc.Dial("127.0.0.1:9090", grpc.WithInsecure())
	client := pb.NewHelloWorldServiceClient(conn)
	_, _ = client.SayHelloWorld(context.Background(), &pb.HttpRequest{Greeting: "jiangnenghua", Infos: map[string]string{
		"name": "april",
	}})
}
