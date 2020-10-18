package main

import (
	"context"
	"daily-test/gRPC"
	"log"
	"net"

	"google.golang.org/grpc"
)

type HelloServiceImpl struct{}

func (p *HelloServiceImpl) Hello(
	ctx context.Context, args *gRPC.String,
) (*gRPC.String, error) {
	reply := &gRPC.String{Value: "hello:" + args.GetValue()}
	return reply, nil
}

func main() {

	grpcServer := grpc.NewServer()
	gRPC.RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)
}
