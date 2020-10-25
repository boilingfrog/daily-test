package main

import (
	"context"
	"daily-test/gRPC_cert"
	"fmt"
	"log"

	"google.golang.org/grpc/credentials"

	"google.golang.org/grpc"
)

func main() {
	// 带入证书的信息
	creds, err := credentials.NewClientTLSFromFile(
		"./gRPC_cert/cert/server.crt", "localhost",
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

	client := gRPC_cert.NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &gRPC_cert.String{Value: "hello"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.GetValue())
}
