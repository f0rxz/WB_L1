package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("go run main.go <worker_number>")
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n <= 0 {
		fmt.Println("Args should be positive")
		return
	}

	dataCh := make(chan int)

	for i := 1; i <= n; i++ {
		go worker(i, dataCh)
	}

	counter := 1
	for {
		dataCh <- counter
		counter++
	}
}

func worker(id int, dataCh <-chan int) {
	for data := range dataCh {
		fmt.Printf("Worker %d got: %d\n", id, data)
	}
}
