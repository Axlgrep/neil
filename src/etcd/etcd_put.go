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

	// 设置key1的值为value1
	key1, value1 := "key1", "value1"
	if resp, err := cli.Put(context.TODO(), key1, value1); err != nil {
		log.Fatal(err)
	} else {
		log.Println(resp)
	}

	// 设置key1的值为value2, 并且返回前一个值
	value2 := "value2"
	if resp, err := cli.Put(context.TODO(), key1, value2, clientv3.WithPrevKV()); err != nil {
		log.Fatal(err)
	} else {
		log.Println(resp)
	}

}
