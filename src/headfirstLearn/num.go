package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	for i := 0; i < 20; i++ {
		go execute()
	}
	fmt.Println(runtime.NumGoroutine())
}
func execute() {
	time.Sleep(time.Minute * 100)
}
