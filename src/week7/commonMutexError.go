// Shared resources (usually memory) should not be accessed by multiple Goroutines at the same time
// This section of code which modifies shared resources is called critical section.
// Mutex (MUTual EXlusion), enables the creation of crital sections
// goroutines need to gain the lock before the continue, if another goroutine currectly has the lock 
// the current goroutine is blocked until it becomes available
package main
import (
    "fmt"
    "sync"
)
var (
    mutex   sync.Mutex
    balance int
)
func init() {
    balance = 5000
}
func deposit(value int, wg *sync.WaitGroup) { // should be ref
    mutex.Lock()
    fmt.Printf("Depositing %d : Current balance: %d\n", value, balance)
    balance += value
    mutex.Unlock()
    wg.Done()
}
func withdraw(value int, wg *sync.WaitGroup) { // should be ref
    mutex.Lock()
    fmt.Printf("Withdrawing %d : Current balance: %d\n", value, balance)
    balance -= value
    mutex.Unlock()
    wg.Done()
}
func main() {
    fmt.Println("Go Mutex Example")
	var wg sync.WaitGroup
    for j := 1; j <= 1000; j++ {
       wg.Add(2)
       go withdraw(300, &wg)  // should be ref
       go deposit(400, &wg) // should be ref
    }

    wg.Wait()
    fmt.Printf("New Balance %d\n", balance)
}