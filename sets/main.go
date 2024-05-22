package main

import "fmt"

func setSolutionA(data []int, x int) bool {
	set := map[int]bool{}
	for _, item := range data {
		set[item] = true
	}

	fmt.Println("set: ", set)
	for k := range set {
		fmt.Printf("%v, ", k)
	}
	fmt.Println()
	return set[x]
}

func setSolutionB(data []int, x int) bool {
	set := map[int]struct{}{}
	for _, v := range data {
		set[v] = struct{}{}
	}
	fmt.Println("set: ", set)
	for k := range set {
		fmt.Printf("%v, ", k)
	}
	fmt.Println()
	if _, ok := set[x]; ok {
		return true
	} else {
		return false
	}
}

func main() {
	data := []int{1, 3, 2, 4, 5, 2, 3, 4, 5}
	fmt.Println("data: ", data)
	fmt.Println("---- solution A ----")
	fmt.Println("found 1:", setSolutionA(data, 1))
	fmt.Println("---- solution B ----")
	fmt.Println("found 10:", setSolutionB(data, 10))
}
