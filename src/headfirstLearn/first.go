package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {

	sq := math.Sqrt(25)
	fmt.Println(strings.Title("Title Page"))
	fmt.Println("Head First Version", "Second Version", sq)
	fmt.Println('A', '\t', 'C')
	fmt.Println(reflect.TypeOf(96.9))
	var firstName string = "Yusuf"
	fmt.Println(firstName)
	var fName, LName string = "Ysuyuf", "Danish"
	fmt.Println(fName, LName)
}
