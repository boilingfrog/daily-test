package main

import (
	"context"
	"daily-test/gRPC_advanced"
	"fmt"
	"log"

	"google.golang.org/grpc/credentials"

	"google.golang.org/grpc"
)

func main() {
	creds, err := credentials.NewClientTLSFromFile(
		"/Users/yj/goWork/daily-test/gRPC_advanced/cert/server.crt", "localhost",
	)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := grpc.Dial("localhost:1234",
		grpc.WithTransportCredentials(creds),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := gRPC_advanced.NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &gRPC_advanced.String{Value: "hello"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.GetValue())

	fmt.Println(grpc.ServiceInfo{})
}
