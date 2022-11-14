package main

import "fmt"

var n [2]int
var m [2]int

func main() {
	n[0] = 10
	n[1] = 20
	m = [2]int{5, 6}
	calc()
	calc1()

	// you cant do statements outside function
	// now calc does not know about n and m so you have to pass them as parameters
}

var q [4]int

func calc() {
	//var finalResult [2]int
	for i := 0; i < 2; i++ {
		//create array
		for j := 0; j < 2; j++ {

			q[i] = n[i] * m[j]
			fmt.Println(q[i])

		}

	}

}

func calc1() {
	for _, result := range m {

		for _, result1 := range n {

			fmt.Println(result * result1)
		}
	}
}
