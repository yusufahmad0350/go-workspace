package main
import "fmt"
func main() {
    ch := make(chan int, 3)
    ch <- 2
    process(ch)
    fmt.Println()
}
func process(ch <-chan int) {
    s := <-ch
    fmt.Println(s)
    //ch <- 2
}
