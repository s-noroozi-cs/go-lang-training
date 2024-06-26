package main

import (
	"fmt"
)

type Stack[T comparable] struct {
	vals []T
}

func (s *Stack[T]) Push(val T) {
	s.vals = append(s.vals, val)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.vals) == 0 {
		var zero T
		return zero, false
	}
	top := s.vals[len(s.vals)-1]
	s.vals = s.vals[:len(s.vals)-1]
	return top, true
}

func (s Stack[T]) Contains(val T) bool {
	for _, v := range s.vals {
		if v == val {
			return true
		}
	}
	return false
}

func main() {
	var intStack Stack[int]

	intStack.Push(10)
	intStack.Push(20)
	intStack.Push(30)

	fmt.Println("stack content: ", intStack)

	v, _ := intStack.Pop()
	fmt.Println("top element: ", v)

	fmt.Println("containt 10: ", intStack.Contains(10))
	fmt.Println("containt 5: ", intStack.Contains(5))

}
