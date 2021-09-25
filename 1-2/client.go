package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func producerConsumer() {
	size := 10
	pipe := make(chan int, size)

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)

	go func() {
		defer func() {
			wg.Done()
			cancel()
			fmt.Println("go producer exit")
		}()

		for i := 0; i < size; i++ {
			select {
			case <-time.After(time.Second):
				fmt.Println("send: ", i+1)
				pipe <- i + 1
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
			fmt.Println("go consumer exit")
		}()

		for {
			select {
			case <-ctx.Done():
				return
			case v := <-pipe:
				fmt.Println("recv: ", v)
			}
		}
	}()

	wg.Wait()
	fmt.Println("main exit.")
}
