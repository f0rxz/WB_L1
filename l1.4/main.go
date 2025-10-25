package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(ctx context.Context, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("worker %d: shutdown\n", id)
			return
		default:
			fmt.Printf("worker %d: working...\n", id)
			time.Sleep(time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	// ловим SIGINT/SIGTERM
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	numWorkers := 10
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(ctx, i, wg)
	}

	<-sigCh
	fmt.Println("\nGot interrupt signal, shutting down...")
	cancel()

	wg.Wait()
	fmt.Println("All workers stopped, exit.")
}
