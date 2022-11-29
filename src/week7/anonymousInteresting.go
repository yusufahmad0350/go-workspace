package main
import (
    "fmt"
    "time"
)
func main() {
     var x string = "Hello World"
     go func() {
         time.Sleep(1 * time.Second)
         // N.B. I can access x in this function even though its not declared in this function    
         fmt.Println("A function with no name and not even a reference",x)
     }()
     x = "Interesting!!"
     time.Sleep(3 * time.Second)
	 
}
