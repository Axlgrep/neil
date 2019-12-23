package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"log"
	"sync"
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

	var w2 sync.WaitGroup
	w2.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			defer w2.Done()
			resp, err := cli.Put(context.TODO(), "keyincr", "", clientv3.WithPrevKV())
			if err != nil {
				fmt.Println(err)
			} else {
				if resp.PrevKv != nil {
					fmt.Println(resp.PrevKv.Version)
				}
			}
		}()
	}

	w2.Wait()
	if resp, err := cli.Get(context.TODO(), "keyincr"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp.Kvs[0].Version)
	}
}
