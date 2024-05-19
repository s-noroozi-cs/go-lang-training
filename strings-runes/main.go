package main

import (
	"fmt"
	"unicode/utf8"
)

func print(str string) {
	fmt.Println("-------------------------------")

	fmt.Printf("value: %v, len: %v\n", str, len(str))

	fmt.Println("")

	for i := 0; i < len(str); i++ {
		fmt.Printf("%x ", str[i])
	}

	fmt.Printf("\n\n")

	fmt.Println("Rune count:", utf8.RuneCountInString(str))

	fmt.Println()
	for idx, runeValue := range str {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	fmt.Println("\nUsing DecodeRuneInString")

	for i, w := 0, 0; i < len(str); i += w {
		runeValue, width := utf8.DecodeRuneInString(str[i:])
		w = width
		if runeValue == 'H' {
			fmt.Println("Found H")
		} else if runeValue == 'و' {
			fmt.Println("Found و")
		}

	}
}

func main() {
	print("درود") // persian hello

	print("Hello") // english hello
}
