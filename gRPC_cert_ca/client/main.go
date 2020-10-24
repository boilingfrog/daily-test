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
	cert, err := tls.LoadX509KeyPair("./gRPC_cert_ca/cert/client/client.pem", "./gRPC_cert_ca/cert/client/client.key")
	if err != nil {
		log.Fatalf("tls.LoadX509KeyPair err: %v", err)
	}

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("./gRPC_cert_ca/cert/ca.pem")
	if err != nil {
		log.Fatalf("ioutil.ReadFile err: %v", err)
	}

	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatalf("certPool.AppendCertsFromPEM err")
	}

	c := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName:   "localhost",
		RootCAs:      certPool,
	})

	conn, err := grpc.Dial(":1234",
		grpc.WithTransportCredentials(c),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := gRPC_cert_ca.NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &gRPC_cert_ca.String{Value: "hello"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.GetValue())

}
