package main

import (
	"fmt"
	"unicode/utf8"
)

func print(str string) {
	fmt.Println("-------------------------------")

	fmt.Printf("value: %v, len: %v\n", str, len(str))

	for i := 0; i < len(str); i++ {
		fmt.Printf("%x ", str[i])
	}
	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(str))

	for idx, runeValue := range str {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

}

func main() {
	print("درود") // persian hello

	print("Hello") // english hello
}
