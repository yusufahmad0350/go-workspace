package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}
func main() {
	messages := make(chan person)
	s := person{name: "Sean", age: 50}
	go func() { messages <- s }()
	p := <-messages
	fmt.Println(p.name)
	fmt.Println(p.age)
}
