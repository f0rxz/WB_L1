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
	if i < 1 || i > 64 {
		fmt.Println("Error: i must be between 1 and 64")
		return
	}

	fmt.Printf("%b\n", number^(1<<(i-1)))
}
