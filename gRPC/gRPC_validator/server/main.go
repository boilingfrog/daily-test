package main

import (
	"context"
	"daily-test/gRPC/gRPC_validator"
	"log"
	"net"

	"google.golang.org/grpc"
)

type HelloServiceImpl struct{}

func (p *HelloServiceImpl) Hello(
	ctx context.Context, args *gRPC_validator.String,
) (*gRPC_validator.String, error) {
	reply := &gRPC_validator.String{Value: "hello:" + args.GetValue()}
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
