package main

import (
	"fmt"
	"sync"
)

var x = 0

// Note pass by reference
func increment(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	defer m.Unlock()
	x = x + 1
	wg.Done()
}
func main() {
	var m sync.Mutex
	var w sync.WaitGroup
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w, &m)
	}
	w.Wait()
	fmt.Println("final value of x", x)
}
