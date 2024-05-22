package main

import "fmt"

func main() {
	x := 10
	fmt.Println("x is ", x)
	if x > 0 {
		x := 5
		fmt.Println("x is ", x)
	}
	fmt.Println("x is ", x)

	if n := 6; n < 0 {
		fmt.Println(n, " is negative")
	} else {
		fmt.Println(n, " is positive")
	}

}
