package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"daily-test/gRPC/gRPC_cert_ca"
	"io/ioutil"
	"log"
	"net"

	"google.golang.org/grpc/credentials"

	"google.golang.org/grpc"
)

type HelloServiceImpl struct{}

func (p *HelloServiceImpl) Hello(
	ctx context.Context, args *gRPC_cert_ca.String,
) (*gRPC_cert_ca.String, error) {
	reply := &gRPC_cert_ca.String{Value: "hello:" + args.GetValue()}
	return reply, nil
}

func main() {

	cert, err := tls.LoadX509KeyPair("./gRPC/gRPC_cert_ca/cert/server/server.pem", "./gRPC/gRPC_cert_ca/cert/server/server.key")
	if err != nil {
		log.Fatalf("tls.LoadX509KeyPair err: %v", err)
	}

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("./gRPC/gRPC_cert_ca/cert/ca.pem")
	if err != nil {
		log.Fatalf("ioutil.ReadFile err: %v", err)
	}

	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatalf("certPool.AppendCertsFromPEM err")
	}

	c := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	})

	grpcServer := grpc.NewServer(grpc.Creds(c))
	gRPC_cert_ca.RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)
}
