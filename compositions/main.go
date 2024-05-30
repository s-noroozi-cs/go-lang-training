package main

import "fmt"

type Employee struct {
	Name string
	ID   string
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

type Manager struct {
	Employee
	Reports []Employee
}

func main() {
	m := Manager{
		Employee: Employee{
			Name: "Saeid Noroozi",
			ID:   "123456",
		},
		Reports: []Employee{},
	}

	fmt.Println("-------- composition - embeded feilds --------")
	fmt.Println(m.ID)
	fmt.Println("-------- composition - method call --------")
	fmt.Println(m.Description())
}
