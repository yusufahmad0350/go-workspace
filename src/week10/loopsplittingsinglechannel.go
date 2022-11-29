/*
	goroutine exercise

Using goroutines\channels write a program that adds the numbers from 1 to some number n using a x number of goroutines
The number of goroutines (x) should be based a loop iteration no greater than 100.
So adding the numbers from 1 to 500
would consist of 5 go routines
goroutine 1 , 1-99
goroutine 2 , 100-199
goroutine 3 , 200-299
goroutine 4 , 300-399
goroutine 5 , 400-500
Besides the use of goroutine make sure that you implement both channels to pass the information and use a select clause to gather your data.
*/
package main

import (
	"fmt"
)

func main() {
	var inVal int
	fmt.Print("Enter number of numbers to generate: ")
	_, err := fmt.Scanf("%d", &inVal)
	if err != nil {
		fmt.Println(err)
	}
	var sum = 0
	for i := 1; i <= inVal; i++ { // assigning 1 to i
		sum += i // sum = sum + i
	}
	fmt.Printf("old way %d \n ", sum)
	var numGoroutine int = (inVal / 100) + 1
	fmt.Printf("I will generate %d \n ", numGoroutine)
	var bottom int = 0
	var top int
	var jump = inVal % 100

	channel := make(chan int, numGoroutine)
	for i := 0; i < numGoroutine; i++ {
		top = (i * 100) + jump
		randomWait(channel, bottom, top)
		bottom = top + 1
	}
	// receive values
	var total = 0
	for i := 0; i < numGoroutine; i++ {
		msg1 := <-channel
		fmt.Printf("received %d from channel \n", msg1)
		total = total + msg1
	}
	fmt.Printf("Total is %d ", total)
}

// Returns channel
func randomWait(c chan int, inFrom int, inTo int) {
	go func() {
		fmt.Printf("I will count from %d to %d \n", inFrom, inTo)
		var sum = 0
		for i := inFrom; i <= inTo; i++ { // assigning 1 to i
			sum += i // sum = sum + i
		}
		c <- sum
	}()
	return // Return the channel to the caller.
}
