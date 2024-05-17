package main

import (
	"fmt"
)

func setZeroByVal(number int) {
	number = 0
}

func setZeroByRef(number *int) {
	*number = 0
}

func main() {
	var number int = 10

	fmt.Printf("number: %v\n", number)

	setZeroByVal(number)
	fmt.Printf("call set zero by val, number: %v\n", number)

	setZeroByRef(&number)
	fmt.Printf("call set zero by ref, number: %v\n", number)

}
