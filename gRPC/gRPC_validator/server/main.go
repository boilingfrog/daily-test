package main

import (
	"context"
	"daily-test/gRPC/gRPC_validator"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type HelloServiceImpl struct{}

func (p *HelloServiceImpl) Hello(ctx context.Context, args *gRPC_validator.RequestInfo) (*gRPC_validator.String, error) {
	if err := args.Validate(); err != nil {
		log.Fatalf("Validate err: %v", err)
	}
	reply := &gRPC_validator.String{Value: fmt.Sprintf("你好：%s,今年：%d岁了", args.GetName(), args.GetAge())}
	return reply, nil
}

func main() {

	grpcServer := grpc.NewServer()
	gRPC_validator.RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)
}
