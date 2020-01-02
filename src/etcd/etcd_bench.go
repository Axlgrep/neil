package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"log"
	"strconv"
	"sync"
	"time"
)

var (
	dialTimeout    = 5 * time.Second
	requestTimeout = 2 * time.Second
	endpoints      = []string{"localhost:2379"}
)

func main() {

	var w2 sync.WaitGroup
	w2.Add(250)
	fmt.Println("start time: ", time.Now())

	for i := 0; i < 250; i++ {
		go func(i int) {
			cli, err := clientv3.New(clientv3.Config{
				Endpoints:   endpoints,
				DialTimeout: dialTimeout,
			})
			if err != nil {
				log.Fatal(err)
				return
			}
			defer cli.Close()

			for idx := 0; idx < 2000; idx++ {
				key1, value1 := "key:"+strconv.Itoa(i)+":"+strconv.Itoa(idx), "value"
				if _, err := cli.Put(context.TODO(), key1, value1); err != nil {
					log.Fatal(err)
				}
			}
			fmt.Printf("gorotine: %d finish\n", i)
			w2.Done()
		}(i)
	}
	w2.Wait()
	fmt.Println("end time: ", time.Now())
}
