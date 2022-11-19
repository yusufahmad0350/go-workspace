package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("number of goroutines A ", runtime.NumGoroutine())
	var mem runtime.MemStats
	fmt.Println("memory baseline...")
	runtime.ReadMemStats(&mem)
	fmt.Println("mem.Alloc:", mem.Alloc)
	fmt.Println("mem.TotalAlloc:", mem.TotalAlloc)
	fmt.Println("mem.HeapAlloc:", mem.HeapAlloc)
	fmt.Println("mem.NumGC:", mem.NumGC)
	leak()
	leak()
	leak()
	fmt.Scanln() // wait for Enter Key
	runtime.ReadMemStats(&mem)
	fmt.Println(mem.Alloc)
	fmt.Println(mem.TotalAlloc)
	fmt.Println(mem.HeapAlloc)
	fmt.Println(mem.HeapSys)
	fmt.Println("number of goroutines C", runtime.NumGoroutine())
}

// leak is a buggy function. It launches a goroutine that
// blocks receiving from a channel. Nothing will ever be
// sent on that channel and the channel is never closed so
// that goroutine will be blocked forever.
func leak() {
	// Creating an array
	arr := [7]string{"This", "is", "an", "example", "of", "Go", "language"}
	ch := make(chan int)
	go func() {
		val := <-ch
		fmt.Println("We received a value:", val)
		// Display array
		fmt.Println("Array:", arr)
	}()
}
