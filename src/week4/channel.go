package main
import "fmt"
func main() {
    ch := make(chan int, 3)
    process(ch)
}
func process(ch chan int) {
    ch <- 2
	ch <- 4
    s := <-ch
	r:= <-ch
    fmt.Println(s)
	fmt.Println(r)
}
