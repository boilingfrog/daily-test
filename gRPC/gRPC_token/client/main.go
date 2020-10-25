package main

import (
	"context"
	"daily-test/gRPC/gRPC_token"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

type Authentication struct {
	User     string
	Password string
}

// 返回认证需要的必要信息
func (a *Authentication) GetRequestMetadata(context.Context, ...string) (
	map[string]string, error,
) {
	return map[string]string{"user": a.User, "password": a.Password}, nil
}

// 表示是否要求底层使用安全链接,测试的代码就是使用了false
func (a *Authentication) RequireTransportSecurity() bool {
	return false
}

func main() {
	// 初始化账户，密码
	auth := Authentication{
		User:     "liz",
		Password: "123456",
	}

	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure(),
		grpc.WithPerRPCCredentials(&auth),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := gRPC_token.NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &gRPC_token.String{Value: "hello"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.GetValue())
}
