package main

import "fmt"

func main() {
	var j = 15.5
	var txt = "Hello World!"

	fmt.Printf("%v\n", j)
	fmt.Printf("%#v\n", j)
	fmt.Printf("%v%%\n", j)
	fmt.Printf("%T\n", j)

	fmt.Printf("%v\n", txt)
	fmt.Printf("%#v\n", txt)
	fmt.Printf("%T\n", txt)

	var i int = 14433322

	fmt.Printf("%b\n", i)
	fmt.Printf("%d\n", i)
	fmt.Printf("%+d\n", i)
	fmt.Printf("%o\n", i)
	fmt.Printf("%O\n", i)
	fmt.Printf("%x\n", i)
	fmt.Printf("%X\n", i)
	fmt.Printf("%#x\n", i)
	fmt.Printf("%4d\n", i)
	fmt.Printf("%-4d\n", i)
	fmt.Printf("%04d\n", i)
}
