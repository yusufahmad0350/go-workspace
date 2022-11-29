package main

import (
	"fmt"
	"sync"
)

var (
	mutex sync.Mutex
	x     int
)

func increment(wg *sync.WaitGroup) {
	mutex.Lock()
	//mutex.Unlock()
	x = x + 1
	mutex.Unlock()
	wg.Done()
}
func main() {
	var w sync.WaitGroup
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w)
	}
	w.Wait()
	fmt.Println("final value of x", x)
}
