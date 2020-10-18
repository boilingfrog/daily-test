package main

import (
	"context"
	"daily-test/gRPC_advanced"
	"log"
	"net"

	"google.golang.org/grpc/credentials"

	"google.golang.org/grpc"
)

type HelloServiceImpl struct{}

func (p *HelloServiceImpl) Hello(
	ctx context.Context, args *gRPC_advanced.String,
) (*gRPC_advanced.String, error) {
	reply := &gRPC_advanced.String{Value: "hello:" + args.GetValue()}
	return reply, nil
}

func main() {

	creds, err := credentials.NewServerTLSFromFile("/Users/yj/goWork/daily-test/gRPC_advanced/cert/server.crt", "/Users/yj/goWork/daily-test/gRPC_advanced/cert/server.key")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer(grpc.Creds(creds))
	gRPC_advanced.RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)
}
