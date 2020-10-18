package main

import (
	"context"
	"daily-test/gRPC_stream"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

type HelloServiceImpl struct{}

func (p *HelloServiceImpl) Hello(
	ctx context.Context, args *gRPC_stream.String,
) (*gRPC_stream.String, error) {
	reply := &gRPC_stream.String{Value: "hello:" + args.GetValue()}
	return reply, nil
}

func (p *HelloServiceImpl) Channel(stream gRPC_stream.HelloService_ChannelServer) error {
	for {
		args, err := stream.Recv()
		if err != nil {
			// io.EOF表示客户端流关闭
			if err == io.EOF {
				return nil
			}
			return err
		}

		reply := &gRPC_stream.String{Value: "hello:" + args.GetValue()}

		err = stream.Send(reply)
		if err != nil {
			return err
		}
	}
}

func main() {

	grpcServer := grpc.NewServer()
	gRPC_stream.RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)
}
