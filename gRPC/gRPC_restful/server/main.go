package main

import (
	"context"
	"daily-test/gRPC/gRPC_restful"
	"log"
	"net"

	"google.golang.org/grpc"
)

type HelloServiceImpl struct{}

func (p *HelloServiceImpl) GetMes(ctx context.Context, args *gRPC_restful.StringMessage) (*gRPC_restful.StringMessage, error) {
	reply := &gRPC_restful.StringMessage{Value: args.Value}
	return reply, nil
}

func (p *HelloServiceImpl) PostMes(ctx context.Context, args *gRPC_restful.StringMessage) (*gRPC_restful.StringMessage, error) {
	reply := &gRPC_restful.StringMessage{Value: args.Value + "post"}
	return reply, nil
}

func main() {
	grpcServer := grpc.NewServer()
	gRPC_restful.RegisterRestServiceServer(grpcServer, new(HelloServiceImpl))

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)
}
