package main

import (
	"fmt"
)

func setZeroByVal(iVal int) {
	iVal = 0
	fmt.Printf("inside setZeroByVal func, ival: %v, type: %T\n", iVal, iVal)
}

func setZeroByRef(iPointer *int) {
	*iPointer = 0
	fmt.Printf("inside setZeroByRef func, ival: %v, type: %T\n", iPointer, iPointer)
}

func main() {
	var number int = 10

	fmt.Printf("number: %v\n", number)

	fmt.Println("---------------------------")

	setZeroByVal(number)
	fmt.Printf("call set zero by val, number: %v\n", number)

	fmt.Println("---------------------------")

	setZeroByRef(&number)
	fmt.Printf("call set zero by ref, number: %v\n", number)

	fmt.Println("---------------------------")

	fmt.Println("pointer:", &number)

}
