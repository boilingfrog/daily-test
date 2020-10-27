package main

import (
	"context"
	"daily-test/gRPC/gRPC_grpcurl"
	"log"
	"net"

	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
)

type HelloServiceImpl struct{}

func (p *HelloServiceImpl) Hello(ctx context.Context, args *gRPC_grpcurl.String) (*gRPC_grpcurl.String, error) {
	reply := &gRPC_grpcurl.String{Value: "hello:" + args.GetValue()}
	return reply, nil
}

func main() {

	grpcServer := grpc.NewServer()
	gRPC_grpcurl.RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))

	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)
}
