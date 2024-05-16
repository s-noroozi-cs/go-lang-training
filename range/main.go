package main

import "fmt"

func main() {

	fmt.Println("---------------------------")

	nums := []int{2, 3, 4}

	for i, v := range nums {
		fmt.Printf("index: %v, value: %v\n", i, v)
	}

	fmt.Println("---------------------------")

	sum := 0
	for _, v := range nums {
		sum += v
	}
	fmt.Printf("sums: %v\n", sum)

	fmt.Println("---------------------------")

	kv := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kv {
		fmt.Printf("key: %v, value: %v\n", k, v)
	}

}
