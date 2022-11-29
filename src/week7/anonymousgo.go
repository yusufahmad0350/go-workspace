package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("A function with no name and not even a reference")
	}()
	time.Sleep(1 * time.Second)
}
