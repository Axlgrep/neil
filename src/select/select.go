package main

import (
	"fmt"
	"strconv"
	"time"
)

func processfun1(chan1 chan string) {
	time.Sleep(time.Duration(2) * time.Second)
	for i := 0; i < 10; i++ {
		chan1 <- "processfun1: " + strconv.Itoa(i)
	}
}

func processfun2(chan2 chan string) {
	for i := 0; i < 10; i++ {
		chan2 <- "processfun2: " + strconv.Itoa(i)
	}
}

func main() {
	chan1 := make(chan string)
	chan2 := make(chan string)

	go processfun1(chan1)
	go processfun2(chan2)
	for {
		select {
		case msg1, ok := <-chan1:
			fmt.Println("msg1: ", msg1, " ok: ", ok)
		case msg2, ok := <-chan2:
			fmt.Println("msg2: ", msg2, " ok: ", ok)
		}
	}
}
