package main

import (
	"context"
	"daily-test/gRPC/gRPC_token"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"

	"google.golang.org/grpc"
)

type Authentication struct {
	User     string
	Password string
}

type grpcServer struct {
	auth *Authentication
}

func (p *grpcServer) Hello(
	ctx context.Context, args *gRPC_token.String,
) (*gRPC_token.String, error) {
	// 初始化信息，测试用的
	p.Init()
	// 检验
	if err := p.auth.Auth(ctx); err != nil {
		return nil, err
	}

	reply := &gRPC_token.String{Value: "hello:" + args.GetValue()}
	return reply, nil
}

// 测试使用
func (p *grpcServer) Init() {
	p.auth = &Authentication{
		User:     "liz",
		Password: "123456",
	}
}

// 认证
func (a *Authentication) Auth(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return fmt.Errorf("missing credentials")
	}

	var appid string
	var appkey string

	if val, ok := md["user"]; ok {
		appid = val[0]
	}
	if val, ok := md["password"]; ok {
		appkey = val[0]
	}

	if appid != a.User || appkey != a.Password {
		return grpc.Errorf(codes.Unauthenticated, "invalid token")
	}

	return nil
}

func main() {

	grpcServer1 := grpc.NewServer()
	gRPC_token.RegisterHelloServiceServer(grpcServer1, new(grpcServer))

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer1.Serve(lis)
}
