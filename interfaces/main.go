package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct{}

func (dog Dog) Speak() string {
	return "Woof!"
}

type Cat struct{}

func (cat Cat) Speak() string {
	return "Meow!"
}

type JavaDeveloper struct{}

func (jd JavaDeveloper) Speak() string {
	return "Design Patterns!"
}

func main() {
	animals := []Animal{Dog{}, Cat{}, JavaDeveloper{}}

	for _, animal := range animals {
		fmt.Println("------------------------")
		fmt.Printf("Type: %T, Speak: %v\n", animal, animal.Speak())
		fmt.Println("")
	}
}
