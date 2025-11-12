package main

import (
	"fmt"
	"math"
)

func bucket(t float64) int {
	q := t / 10
	if t >= 0 {
		return int(math.Floor(q)) * 10
	}
	return int(math.Ceil(q)) * 10
}

func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groups := make(map[int][]float64)

	for _, t := range temps {
		k := bucket(t)
		groups[k] = append(groups[k], t)
	}

	keys := make([]int, 0, len(groups))
	for k := range groups {
		keys = append(keys, k)
	}

	for _, k := range keys {
		fmt.Printf("%d: %v\n", k, groups[k])
	}
}
