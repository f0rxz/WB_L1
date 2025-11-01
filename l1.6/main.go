package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("goroutine 1 stopped by context")
				return
			default:
				fmt.Println("goroutine 1")
				time.Sleep(time.Millisecond * 200)
			}
		}
	}(ctx)

	go func() {
		i := 0
		for {
			fmt.Println("goroutine 2")
			time.Sleep(time.Millisecond * 200)
			i++
			if i == 5 {
				fmt.Println("goroutine 2 stopped by condition")
				return
			}
		}
	}()

	done := make(chan struct{})

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("goroutine 3 stopped by channel")
				return
			default:
				fmt.Println("goroutine 3")
				time.Sleep(time.Millisecond * 200)
			}
		}
	}()

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("goroutine 4")
		}
		fmt.Println("goroutine 4 stopped by runtime")
		runtime.Goexit()
	}()

	time.Sleep(time.Second * 1)
	cancel()
	close(done)
	time.Sleep(time.Second)
	fmt.Println("main goroutine exit")

}
