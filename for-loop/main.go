package main

import (
	"fmt"
)

func main() {
	i := 0
	for i < 3 {
		fmt.Printf("for loop A, %v\n", i)
		i = i + 1
	}

	fmt.Println("--------------------------")

	for j := 0; j < 3; j++ {
		fmt.Printf("for loop B, %v\n", j)
	}

	fmt.Println("--------------------------")

	for i := range 3 {
		fmt.Printf("for loop C, %v\n", i)
	}

	fmt.Println("--------------------------")

	for {
		fmt.Printf("for loop D, break!\n")
		break
	}

	fmt.Println("--------------------------")

	for i := range 10 {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("for loop E, continue, odd number range 10,  %v\n", i)
	}

}
