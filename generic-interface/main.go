package main

import (
	"fmt"
	"math"
)

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

type Point2D struct {
	X, Y int
}

func (p Point2D) String() string {
	return fmt.Sprintf("{%d, %d}", p.X, p.Y)
}

func (p Point2D) Diff(from Point2D) float64 {
	x := p.X - from.X
	y := p.Y - from.Y
	return math.Sqrt(float64(x*x) + float64(y*y))
}

type Point3D struct {
	X, Y, Z int
}

func (p Point3D) String() string {
	return fmt.Sprintf("{%d, %d, %d}", p.X, p.Y, p.Z)
}

func (p Point3D) Diff(from Point3D) float64 {
	x := p.X - from.X
	y := p.Y - from.Y
	z := p.Z - from.Z
	return math.Sqrt(float64(x*x) + float64(y*y) + float64(z*z))
}

func main() {
	pair2Da := Pair[Point2D]{Point2D{1, 1}, Point2D{5, 5}}
	pair2Db := Pair[Point2D]{Point2D{10, 10}, Point2D{15, 5}}
	closer2D := FindCloser(pair2Da, pair2Db)
	fmt.Printf("Point A: %v\nPoint B: %v\nCloser Point: %v\n", pair2Da, pair2Db, closer2D)

	fmt.Println("-----------------------------------------")

	pair3Da := Pair[Point3D]{Point3D{1, 1, 10}, Point3D{5, 5, 0}}
	pair3Db := Pair[Point3D]{Point3D{10, 10, 10}, Point3D{11, 5, 0}}
	closer3D := FindCloser(pair3Da, pair3Db)
	fmt.Printf("Point A: %v\nPoint B: %v\nCloser Point: %v\n", pair3Da, pair3Db, closer3D)

}
