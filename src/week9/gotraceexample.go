package main

//  to view the trace file "go tool trace trace.out"
import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
)

func main() {
	fmt.Println("number of goroutines A", runtime.NumGoroutine())
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()
	leak()
	leak()
	leak()
	leak()
	fmt.Scanln() // wait for Enter Key
	fmt.Println("number of goroutines B", runtime.NumGoroutine())
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
