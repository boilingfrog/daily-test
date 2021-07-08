package main

import (
	"fmt"
	"log"

	clientv3 "go.etcd.io/etcd/client/v3"
	recipe "go.etcd.io/etcd/client/v3/experimental/recipes"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"localhost:2379"},
	})
	if err != nil {
		log.Fatalf("error New (%v)", err)
	}

	done := make(chan struct{})
	go func() {
		defer func() {
			done <- struct{}{}
		}()
		q := recipe.NewQueue(cli, "testq")
		for i := 0; i < 5; i++ {
			if err := q.Enqueue(fmt.Sprintf("%d", i)); err != nil {
				log.Fatalf("error enqueuing (%v)", err)
			}
		}
	}()

	q := recipe.NewQueue(cli, "testq")
	for i := 0; i < 5; i++ {
		s, err := q.Dequeue()
		if err != nil {
			log.Fatalf("error dequeueing (%v)", err)
		}
		if s != fmt.Sprintf("%d", i) {
			log.Fatalf("+++++++expected dequeue value %v, got %v", s, i)
		}
		fmt.Println(s)
	}
	<-done
}
