package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"log"
	"time"
)

var (
	dialTimeout    = 5 * time.Second
	requestTimeout = 2 * time.Second
	endpoints      = []string{"localhost:2379"}
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: dialTimeout,
	})

	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	go func() {
		rch := cli.Watch(context.Background(), "", clientv3.WithPrefix())
		for wresp := range rch {
			for _, ev := range wresp.Events {
				fmt.Printf("Watch: %s %q: %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			}
		}
	}()

	key1, value1 := "key1", "value1"
	if resp, err := cli.Put(context.TODO(), key1, value1); err != nil {
		log.Fatal(err)
	} else {
		log.Println(resp)
	}
	time.Sleep(time.Duration(2) * time.Second)
}
