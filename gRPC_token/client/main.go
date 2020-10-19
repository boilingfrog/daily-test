package main

import (
	"context"
	"daily-test/gRPC_token"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func main() {
	auth := gRPC_token.Authentication{
		User:     "liz",
		Password: "123456",
	}

	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure(), grpc.WithPerRPCCredentials(&auth))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := gRPC_token.NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &gRPC_token.String{Value: "hello"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.GetValue())
}
