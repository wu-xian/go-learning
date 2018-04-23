package main

import (
	"fmt"
	"learn/src/grpc-learning/go-client/grpcg"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("go client started")
	data := &grpcg.HelloRequest{
		Name: "hellojaja",
	}
	conn, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	client := grpcg.NewGrpcClient(conn)
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	response, err := client.SayHello(ctx, data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response)
}
