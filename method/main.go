package main

import (
	"fmt"
	"time"
)

type Counter struct {
	total       int
	lastUpdated time.Time
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("total: %d, lastUpdate: %v", c.total, c.lastUpdated)
}

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
	fmt.Println("----- method A -----")
	rect := Rect{width: 10, height: 5}

	fmt.Printf("Rect Area:\t%v\n", rect.Area())
	fmt.Printf("Rect Perim:\t%v\n", rect.Perim())

	fmt.Println("----- method B -----")
	var counter Counter
	fmt.Println(counter.String())
	counter.Increment()
	fmt.Println(counter.String())
}
