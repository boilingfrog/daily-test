package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"daily-test/gRPC/gRPC_web"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"google.golang.org/grpc/credentials"

	"google.golang.org/grpc"
)

type HelloServiceImpl struct{}

func (p *HelloServiceImpl) Hello(
	ctx context.Context, args *gRPC_web.String,
) (*gRPC_web.String, error) {
	reply := &gRPC_web.String{Value: "hello:" + args.GetValue()}
	return reply, nil
}

func main() {

	cert, err := tls.LoadX509KeyPair("./gRPC/gRPC_web/cert/server/server.pem", "./gRPC/gRPC_web/cert/server/server.key")
	if err != nil {
		log.Fatalf("tls.LoadX509KeyPair err: %v", err)
	}

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("./gRPC/gRPC_web/cert/ca.pem")
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

	gRPC_web.RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("hello"))
	})

	http.ListenAndServeTLS(":1234", "./gRPC/gRPC_web/cert/server/server.pem", "./gRPC/gRPC_web/cert/server/server.key",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.ProtoMajor != 2 {
				mux.ServeHTTP(w, r)
				return
			}
			if strings.Contains(
				r.Header.Get("Content-Type"), "application/grpc",
			) {
				grpcServer.ServeHTTP(w, r)
				return
			}

			mux.ServeHTTP(w, r)
			return
		}),
	)
}
