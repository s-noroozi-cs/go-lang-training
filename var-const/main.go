package main

import "fmt"

func main() {
	var i int = 20
	var f float64 = float64(i)
	fmt.Println(i, f)

	const c = 20
	i = c
	f = c
	fmt.Println(i, f)

	var b byte = 0xff
	var si uint8 = 0xff
	var bi uint16 = 0xffff

	fmt.Println(b, si, bi)

	b++
	si++
	bi++

	fmt.Println(b, si, bi)
}
