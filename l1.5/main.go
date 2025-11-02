package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	defer close(ch)
	done := time.After(time.Duration(2 * time.Second))

	go func() {
		i := 0

		for {
			ch <- i
			i++
			time.Sleep(300 * time.Millisecond)
		}
	}()

	for {
		select {
		case <-done:
			fmt.Println("timeout, exit")
			return
		case msg := <-ch:
			fmt.Println("received:", msg)
		}
	}

}
