package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"daily-test/gRPC_cert_ca"
	"fmt"
	"io/ioutil"
	"log"

	"google.golang.org/grpc/credentials"

	"google.golang.org/grpc"
)

func main() {
	// 带入证书的信息
	//creds, err := credentials.NewClientTLSFromFile(
	//	"/Users/yj/goWork/daily-test/gRPC_cert_ca/cert/server.crt", "localhost",
	//)
	//if err != nil {
	//	log.Fatal(err)
	//}

	cert, err := tls.LoadX509KeyPair("/Users/yj/goWork/daily-test/gRPC_cert_ca/cert/client/client.pem", "/Users/yj/goWork/daily-test/gRPC_cert_ca/cert/client/client.key")
	if err != nil {
		log.Fatalf("tls.LoadX509KeyPair err: %v", err)
	}

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("/Users/yj/goWork/daily-test/gRPC_cert_ca/cert/ca.pem")
	if err != nil {
		log.Fatalf("ioutil.ReadFile err: %v", err)
	}

	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatalf("certPool.AppendCertsFromPEM err")
	}

	c := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName:   "gRPC_cert_ca",
		RootCAs:      certPool,
	})

	conn, err := grpc.Dial("localhost:1234",
		grpc.WithTransportCredentials(c),
	)
	if err != nil {
		log.Fatal(err)
		fmt.Println("+++++++++")
	}
	defer conn.Close()

	client := gRPC_cert_ca.NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &gRPC_cert_ca.String{Value: "hello"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.GetValue())
}
