// The len function for a channel gives the number of elements queued in the channel buffer
// The cap function for a channel gives the number of capacity of the channel buffer
package main

import "fmt"

func main() {
	ch := make(chan int, 8)
	ch <- 42
	ch <- 7
	<-ch
	ch <- 64
	<-ch
	ch <- 67
	// number of queued elements = 1 + 1 - 1 + 1 = 2
	fmt.Println(len(ch), cap(ch))
}
