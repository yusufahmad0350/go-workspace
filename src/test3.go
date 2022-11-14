package main

import (
    "fmt"
    "sync"
)

var x = 0


func increment(wg *sync.WaitGroup,mu *sync.Mutex) {
	
	mu.Lock()
    x = x + 1
    wg.Done()
	mu.Unlock()
}

func main() {
	
    var w sync.WaitGroup
	var mu       sync.Mutex
    for i := 0; i < 1000; i++ {
        w.Add(1)
        go increment(&w, &mu)
    }
	

    w.Wait()
    fmt.Println("final value of x", x)
}