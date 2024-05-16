package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("------------------")

	i := 2
	fmt.Print("convert ", i, " is ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	fmt.Println("------------------")

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's the weekend")
	default:
		fmt.Println("it's a weekday")
	}

	fmt.Println("------------------")

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("it's before noon")
	default:
		fmt.Println("it's after noon")
	}

	fmt.Println("------------------")

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Printf("I'm a bool, value: %v\n", i)
		case int:
			fmt.Printf("I'm a int, value: %v\n", i)
		default:
			fmt.Printf("Don't know, value: %v, type %T\n", i, t)
		}
	}

	whatAmI(true)
	whatAmI(12)
	whatAmI("Saeid")
}
