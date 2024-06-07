package main

import (
	"cmp"
	"fmt"
)

type OrderableFunc[T any] func(t1, t2 T) int

type Tree[T any] struct {
	f    OrderableFunc[T]
	root *Node[T]
}

type Node[T any] struct {
	val         T
	left, right *Node[T]
}

func NewTree[T any](f OrderableFunc[T]) *Tree[T] {
	return &Tree[T]{
		f: f,
	}
}

func (t *Tree[T]) Add(v T) {
	t.root = t.root.Add(t.f, v)
}

func (t *Tree[T]) Contains(v T) bool {
	return t.root.Contains(t.f, v)
}

func (n *Node[T]) Add(f OrderableFunc[T], v T) *Node[T] {
	if n == nil {
		return &Node[T]{val: v}
	}
	switch r := f(v, n.val); {
	case r <= -1:
		n.left = n.left.Add(f, v)
	case r >= 1:
		n.right = n.right.Add(f, v)
	}
	return n
}

func (n *Node[T]) Contains(f OrderableFunc[T], v T) bool {
	if n == nil {
		return false
	}
	switch r := f(v, n.val); {
	case r <= -1:
		return n.left.Contains(f, v)
	case r >= 1:
		return n.right.Contains(f, v)
	}
	return true
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("Name: %s, Age: %d", p.Name, p.Age)
}

func (p Person) Order(other Person) int {
	out := cmp.Compare(p.Name, other.Name)
	if out == 0 {
		out = cmp.Compare(p.Age, other.Age)
	}
	return out
}

func main() {
	tree := NewTree(cmp.Compare[int])

	for _, value := range []int{10, 30, 15} {
		fmt.Printf("Adding %v to tree\n", value)
		tree.Add(value)
	}

	fmt.Println("-------------------------------------------")

	for _, value := range []int{15, 40} {
		fmt.Printf("Searching %v in tree, result: %v\n", value, tree.Contains(value))
	}

	fmt.Println("-------------------------------------------")

	personTree := NewTree(Person.Order)
	saeid := Person{"Saeid", 37}
	reza := Person{"Reza", 40}
	davood := Person{"Davood", 43}

	for _, item := range []Person{saeid, reza, davood} {
		fmt.Printf("Adding new person %v to tree\n", item)
		personTree.Add(item)
	}

	fmt.Println("-------------------------------------------")

	for _, item := range []Person{saeid, {"Dadash", 58}} {
		fmt.Printf("Person Tree contain %v: %v\n", item, personTree.Contains(item))
	}
}
