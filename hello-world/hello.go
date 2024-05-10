package main

import "fmt"

func main() {
	var msg = "Hello\n\"world\""
	fmt.Println(msg)
	msg = `Hello 
"world"`
	fmt.Println(msg)
}
