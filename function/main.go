package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func divAndRemainder(num, denom int) (int, int, error) {
	if denom == 0 {
		return 0, 0, errors.New("CAN'T DIVID BY ZERO")
	}
	return num / denom, num % denom, nil
}

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, v := range nums {
		total += v
	}
	fmt.Println(total)
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
func add(a int, b int) int { return a + b }
func sub(a int, b int) int { return a - b }
func mul(a int, b int) int { return a * b }
func div(a int, b int) int { return a / b }

func main() {
	a := 1
	b := 2
	fmt.Printf("adding %v and %v is %v\n", a, b, add(a, b))

	i, j, k, l := calc(a, b)
	fmt.Printf("four main opr(+ - * /)  on %v and %v is: %v %v %v %v\n", a, b, i, j, k, l)

	x := 10
	fmt.Printf("factorial %v is %v\n", x, fact(x))

	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum(nums...)

	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		} else {
			return fib(n-1) + fib(n-2)
		}
	}

	fmt.Printf("fib 10: %v\n", fib(10))

	result, remainder, error := divAndRemainder(5, 3)
	if error != nil {
		fmt.Println(error)
		os.Exit(1)
	} else {
		fmt.Printf("5, 3 --> div: %v, rem: %v\n", result, remainder)
	}

	/*result, remainder, error = divAndRemainder(5, 0)
	if error != nil {
		fmt.Println(error)
		os.Exit(1)
	} else {
		fmt.Printf("5, 0 --> div: %v, rem: %v\n", result, remainder)
	}*/

	var oprMap = map[string]func(int, int) int{
		"+": add,
		"-": sub,
		"*": mul,
		"/": div,
	}

	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "%", "3"},
		{"two", "+", "three"},
		{"2", "^", "3"},
		{"-", "2"},
	}

	for _, exp := range expressions {
		if len(exp) != 3 {
			fmt.Println(exp, " invalida expression.")
			continue
		}
		p1, err := strconv.Atoi(exp[0])
		if err != nil {
			fmt.Println(exp, ", invalid operant 1: ", err)
			continue
		}
		opr := exp[1]
		opFunc, ok := oprMap[opr]
		if !ok {
			fmt.Println(exp, ", unsopported operator: ", opr)
			continue
		}
		p2, err := strconv.Atoi(exp[2])
		if err != nil {
			fmt.Println("invalid operant 2: ", err)
			continue
		}
		fmt.Println(exp, ": ", opFunc(p1, p2))
	}

}
