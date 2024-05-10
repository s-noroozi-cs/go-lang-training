package main

import "fmt"

func main() {
	var msg = "Hello\n\"world\""
	fmt.Println(msg)
	msg = `Hello 
"world"`
	fmt.Println(msg)
	var x int = 10
	var y float64 = 30.2
	var sumf float64 = float64(x) + y
	var sumi int = x + int(y)
	fmt.Println(sumf, sumi)
}
