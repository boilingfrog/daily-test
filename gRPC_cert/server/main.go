package main

import (
	"context"
	"daily-test/gRPC_cert"
	"log"
	"net"

	"google.golang.org/grpc/credentials"

	"google.golang.org/grpc"
)

type HelloServiceImpl struct{}

func (p *HelloServiceImpl) Hello(
	ctx context.Context, args *gRPC_cert.String,
) (*gRPC_cert.String, error) {
	reply := &gRPC_cert.String{Value: "hello:" + args.GetValue()}
	return reply, nil
}

func main() {

	creds, err := credentials.NewServerTLSFromFile("/Users/yj/goWork/daily-test/gRPC_cert/cert/server.crt", "/Users/yj/goWork/daily-test/gRPC_cert/cert/server.key")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer(grpc.Creds(creds))
	gRPC_cert.RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)
}
