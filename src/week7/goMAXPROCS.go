package main
import (
    "runtime"
    "fmt"
)

func main() {
    // NumCPU returns the number of logical CPUs usable by the current process
    fmt.Printf("Number of CPUs is %d\n", runtime.NumCPU())
    fmt.Printf("GOMAXPROCS is %d\n", runtime.GOMAXPROCS(0))
    runtime.GOMAXPROCS(2)
    fmt.Printf("GOMAXPROCS is %d\n", runtime.GOMAXPROCS(0))
}