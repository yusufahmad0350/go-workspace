package main

import "fmt"

func main() {
	a()
	var functions []func() // array of functions
	for i := 0; i < 10; i++ {
		// build an array of functions to call later
		functions = append(functions, func() {
			fmt.Println(i) // just print i
		})
	}
	// now lets call all the function in the array in order
	for _, f := range functions {
		f()
	}
}

//To solve the above issue

func a() {
	var functions []func()
	for i := 0; i < 10; i++ {
		functions = append(functions, func(val int) func() {
			return func() {
				fmt.Println(val)
			}
		}(i)) // pass by value in Go (so a copy is made)
	}
	for _, f := range functions {
		f()
	}
}
