package main

import (
	"context"
	"daily-test/gRPC/gRPC_base"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := gRPC_base.NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &gRPC_base.String{Value: "hello"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.GetValue())
}
