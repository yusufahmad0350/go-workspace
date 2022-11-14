// In computer science, a nondeterministic algorithm is an algorithm that, even for the same input, can exhibit different behaviors on different runs
// We spawn 1000 increment Goroutines from line no. 15 of the \\program above
// Each of these Goroutines run concurrently and \\race condition occurs when trying to increment x is line no. 8 as
// multiple Goroutines try to access the value of x concurrently.
package main

import (
	"fmt"
	"sync"
)

var x = 0

// Note pass by reference
func increment(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	//m.Unlock()  // Without defer, non deterministic
	defer m.Unlock() //deterministic
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
