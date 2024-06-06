package main

import (
	"errors"
	"fmt"
)

type Integer interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		uintptr
}

func divAndRemainder[T Integer](num, denom T) (T, T, error) {
	fmt.Printf("Type num: %T, Value num: %v\n", num, num)

	fmt.Printf("Type denom: %T, Value denom: %v\n", denom, denom)

	if denom == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}

	div, remainder := num/denom, num%denom

	fmt.Printf("Type div: %T, Value div: %v\n", div, div)
	fmt.Printf("Type Remainder: %T, Value Remainder: %v\n", remainder, remainder)

	fmt.Println("----------------------------")

	return div, remainder, nil
}

func main() {
	divAndRemainder(int32(10), int32(4))

	divAndRemainder(int8(10), int8(4))

	divAndRemainder(uint(10), uint(4))

}
