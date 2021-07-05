package discovery

import (
	"context"
	"fmt"
	"log"
	"net"
	"testing"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/resolver"

	"daily-test/etcd/discovery/helloworld"

	"google.golang.org/grpc"
)

var etcdAddrs = []string{"127.0.0.1:2379"}

func TestResolver(t *testing.T) {
	r := NewResolver(etcdAddrs, zap.NewNop())
	resolver.Register(r)

	go newServer(t, ":8081", "1.0.0", 1)
	go newServer(t, ":8082", "1.0.0", 10)

	conn, err := grpc.Dial("etcd:///hello", grpc.WithInsecure(), grpc.WithBalancerName(roundrobin.Name))
	if err != nil {
		t.Fatalf("failed to dial %v", err)
	}
	defer conn.Close()

	c := helloworld.NewGreeterClient(conn)

	go func() {
		time.Sleep(1 * time.Second)
		go newServer(t, ":8083", "1.0.1", 4)
		time.Sleep(1 * time.Second)
		go newServer(t, ":8084", "1.0.0", 5)
	}()

	for i := 0; i < 10; i++ {
		resp, err := c.SayHello(context.Background(), &helloworld.HelloRequest{Name: "abc"})
		if err != nil {
			t.Fatalf("say hello failed %v", err)
		}
		log.Println(resp.Message)
		time.Sleep(100 * time.Millisecond)
	}

	time.Sleep(10 * time.Second)
}

type server struct {
	Port string
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{Message: fmt.Sprintf("Hello From %s", s.Port)}, nil
}

func newServer(t *testing.T, port string, version string, weight int64) {
	register := NewRegister(etcdAddrs, zap.NewNop())
	defer register.Stop()

	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	s := grpc.NewServer()
	helloworld.RegisterGreeterServer(s, &server{Port: port})

	info := Server{
		Name:    "hello",
		Addr:    fmt.Sprintf("127.0.0.1%s", port),
		Version: version,
		Weight:  weight,
	}
	if port == ":8082" {
		go func(r *Register) {
			time.Sleep(10 * time.Second)
			fmt.Println("stop 8082")
			r.Stop()
		}(register)
	}

	register.Register(info, 10)

	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to server %v", err)
	}
}
