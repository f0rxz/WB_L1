package main

import (
	"fmt"
	"sync"
)

func main() {
	data := make(map[int]int)
	var mu sync.Mutex
	wg := sync.WaitGroup{}
	numWrites := 100

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < numWrites; i += 2 {
			mu.Lock()
			data[i] = i
			mu.Unlock()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i < numWrites; i += 2 {
			mu.Lock()
			data[i] = i
			mu.Unlock()
		}
	}()

	wg.Wait()
	fmt.Println(data)
}
