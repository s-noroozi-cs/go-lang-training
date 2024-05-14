package main

import (
	"fmt"
)

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func printPerson(person Person) {
	fmt.Printf("Name: %v, Age: %v, Job: %v, Salary: %v\n", person.name, person.age, person.job, person.salary)
}

func main() {
	var person Person
	person.name = "Saeid"
	person.age = 38
	person.job = "Software Engineer"
	person.salary = 2_000

	printPerson(person)
}
