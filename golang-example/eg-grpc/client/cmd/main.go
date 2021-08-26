package main

import (
	"context"
	"example/eg-grpc/hello"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
)

const (
	address     = "127.0.0.1:50051"
	defaultName = "world"
)

func main() {
	// 设置grpc服务器地址
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := hello.NewGreeterClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &hello.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
