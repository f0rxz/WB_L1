package main

import (
	"fmt"
	"math/rand"
)

func main() {
	numbers := make(chan int)
	results := make(chan int)
	arr := make([]int, 100)

	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(1000)
	}

	go func() {
		defer close(numbers)
		for _, number := range arr {
			numbers <- number
		}

	}()

	go func() {
		defer close(results)
		for number := range numbers {
			square := number * number
			results <- square
		}
	}()

	for number := range results {
		fmt.Println(number)
	}
}
