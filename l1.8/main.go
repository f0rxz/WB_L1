package main

import "fmt"

func main() {
	var number int64
	if _, err := fmt.Scanf("%d", &number); err != nil {
		fmt.Printf("Error scanning number: %v\n", err)
	}
	fmt.Printf("You entered: %d\n", number)
	var i int
	if _, err := fmt.Scanf("%d", &i); err != nil {
		fmt.Printf("Error scanning i: %v\n", err)
	}
	fmt.Printf("You entered: %d\n", i)

	fmt.Printf("%b\n", number^(1<<(i-1)))
}
