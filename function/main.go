package main

import (
	"fmt"
)

func add(a int, b int) int {
	return a + b
}

func calc(a int, b int) (int, int, int, int) {
	return a + b, a - b, a * b, a / b
}

func fact(a int) int {
	if a > 0 {
		return a * fact(a-1)
	} else {
		return 1
	}
}

func main() {
	a := 1
	b := 2
	fmt.Printf("adding %v and %v is %v\n", a, b, add(a, b))

	add, sub, mul, div := calc(a, b)
	fmt.Printf("four main opr(+ - * /)  on %v and %v is: %v %v %v %v\n", a, b, add, sub, mul, div)

	x := 10
	fmt.Printf("factorial %v is %v\n", x, fact(x))
}
