package main

import (
	"fmt"
)

type Rect struct {
	width  int
	height int
}

func (rect Rect) Area() int {
	return rect.width * rect.height
}

func (rect Rect) Perim() int {
	return 2*rect.width + 2*rect.height
}

func main() {
	fmt.Println("----- method -----")
	rect := Rect{width: 10, height: 5}

	fmt.Printf("Rect Area:\t%v\n", rect.Area())
	fmt.Printf("Rect Perim:\t%v\n", rect.Perim())
}
