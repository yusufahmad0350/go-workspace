// The len function for a channel gives the number of elements queued in the channel buffer
// The cap function for a channel gives the number of capacity of the channel buffer
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() { ch <- 42 }()
	time.Sleep(2 * time.Second)
	fmt.Println(len(ch), cap(ch))
	fmt.Println(<-ch)
}
