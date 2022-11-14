package main

import "fmt"

func main() {
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
