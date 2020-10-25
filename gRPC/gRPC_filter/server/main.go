package main

import (
	"context"
	"daily-test/gRPC/gRPC_filter"
	"log"
	"net"

	"google.golang.org/grpc"
)

type HelloServiceImpl struct{}

func (p *HelloServiceImpl) Hello(
	ctx context.Context, args *gRPC_filter.String,
) (*gRPC_filter.String, error) {
	reply := &gRPC_filter.String{Value: "hello:" + args.GetValue()}
	return reply, nil
}

func filter(ctx context.Context,
	req interface{}, info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	log.Println("filter:", info)
	return handler(ctx, req)
}

func main() {

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(filter))
	gRPC_filter.RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)
}
