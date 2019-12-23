package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"log"
	"math/rand"
	"sync"
	"time"
)

var (
	dialTimeout    = 5 * time.Second
	requestTimeout = 2 * time.Second
	endpoints      = []string{"localhost:2379"}
)

func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

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
	w2.Add(10)

	key10 := "setnx"
	rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		go func(i int) {
			time.Sleep(time.Duration(rand.Intn(5)) * time.Millisecond)
			_, err := cli.Txn(context.Background()).
				If(clientv3.Compare(clientv3.CreateRevision(key10), "=", 0)).
				Then(clientv3.OpPut(key10, fmt.Sprintf("%d", i))).
				Commit()
			if err != nil {
				fmt.Println(err)
			}
			w2.Done()
		}(i)
	}

	w2.Wait()
	if resp, err := cli.Get(context.TODO(), key10); err != nil {
		log.Fatal(err)
	} else {
		log.Println(resp)
	}
}
