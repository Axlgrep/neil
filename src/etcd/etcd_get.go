package main

import (
	"context"
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

	key1 := "key1"
	// 读取一个已经存在的值
	if resp, err := cli.Get(context.TODO(), key1); err != nil {
		log.Fatal(err)
	} else {
		log.Println(resp)
	}
	// 读取一个不存在的值
	if resp, err := cli.Get(context.TODO(), key1+key1); err != nil {
		log.Fatal(err)
	} else {
		log.Println(resp)
	}
}
