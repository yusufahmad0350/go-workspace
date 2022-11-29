package main

import "fmt"

func main() {
	func(message string) {
		fmt.Println(message)
	}(" A function with no name and an inline call!")

	aFunc := func() {
		fmt.Println("A function with no name but only a reference")
	}
	aFunc()
}
