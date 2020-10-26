package main

import (
	"context"
	"daily-test/gRPC/gRPC_restful"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
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
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()

	err := gRPC_restful.RegisterRestServiceHandlerFromEndpoint(
		ctx, mux, "localhost:1234",
		[]grpc.DialOption{grpc.WithInsecure()},
	)
	if err != nil {
		log.Fatal(err)
	}

	http.ListenAndServe(":8080", mux)

}
