package main

import "fmt"

type Pair[T fmt.Stringer] struct {
	ValA T
	ValB T
}

type Differ[T any] interface {
	fmt.Stringer
	Diff(T) float64
}

func FindCloser[T Differ[T]](pairA, pairB Pair[T]) Pair[T] {
	d1 := pairA.ValA.Diff(pairA.ValB)
	d2 := pairB.ValA.Diff(pairB.ValB)
	if d1 < d2 {
		return pairA
	}
	return pairB
}

func main() {
	fmt.Println("ok")
}
