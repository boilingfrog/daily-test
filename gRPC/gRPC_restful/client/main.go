package main

import (
	"context"
	"daily-test/gRPC/gRPC_restful"
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

	client := gRPC_restful.NewRestServiceClient(conn)
	reply, err := client.GetMes(context.Background(), &gRPC_restful.StringMessage{Value: "hello"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.GetValue())
}
