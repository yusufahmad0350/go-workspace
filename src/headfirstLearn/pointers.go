package main

import "fmt"

func main() {

	myInt := 4
	myIntpointer := &myInt
	fmt.Println(myIntpointer)
	fmt.Println(*myIntpointer)
	fmt.Println(&myIntpointer)

	*myIntpointer = 8
	fmt.Println(*myIntpointer)
	fmt.Println(myInt)

	var myInt1 int
	myIntPointer1 := &myInt1
	*myIntPointer1 = 42
	fmt.Println(*myIntPointer1)

}
