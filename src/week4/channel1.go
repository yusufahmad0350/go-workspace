package main
import "fmt"
func main() {
    ch := make(chan int, 3)
    process(ch)
    fmt.Println(<-ch)
}
func process(ch chan<- int) {
    ch <- 2
    s := <-ch
}


//Signature of such a bidirectional channel will be like below when passed to a function as an argument.
