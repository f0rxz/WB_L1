package main

import "fmt"

func intersection(a, b []int) []int {
	if len(a) == 0 || len(b) == 0 {
		return nil
	}

	if len(a) > len(b) {
		a, b = b, a
	}

	set := make(map[int]struct{}, len(a))
	for _, v := range a {
		set[v] = struct{}{}
	}

	res := make([]int, 0, len(set))
	for _, v := range b {
		if _, ok := set[v]; ok {
			res = append(res, v)
			delete(set, v)
		}
	}

	return res
}

func main() {
	A := []int{1, 2, 3}
	B := []int{2, 3, 4}
	fmt.Println(intersection(A, B))
}
