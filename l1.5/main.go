package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	done := make(chan struct{})

	go func() {
		defer close(ch)
		i := 0
		for {
			select {
			case <-done:
				return
			case ch <- i:
				i++
				time.Sleep(300 * time.Millisecond)
			}
		}
	}()

	go func() {
		for msg := range ch {
			fmt.Println(msg)
		}
	}()

	time.Sleep(time.Second * 2)
	close(done)
}
