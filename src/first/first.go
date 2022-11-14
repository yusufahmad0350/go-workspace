package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func add(x int, y int) int { //type comes after the variable name
	return (x + y)
}
func sub(a, b int) int { //When two or more consecutive named function parameters share a type, you can omit the type from all but the last.
	return (a - b)
}
func swap(c, d string) (string, string) { //The swap function returns two strings.
	return d, c
}
func split(sum int) (f, g int) { //A return statement without arguments returns the named return values. This is known as a "naked" return.
	f = sum * 4 / 9
	g = sum - f
	return
}

func (xxx int) int() {

	if xxx > 0 {
		fmt.Println(true)
	}

}

var ab, bc, cd bool
var ef, gh, ig int = 1, 2, 3
var i int
var fff float64
var bbb bool
var sss string

var var1, var2, var3 int //The var statement declares a list of variables; as in function argument lists, the type is last.
//var sqroot complex128 = math.Sqrt(-5 + 2i)

func main() {
	fmt.Println("This is my first progra")
	fmt.Println(time.Now())
	fmt.Println(rand.Intn(10))
	fmt.Println(math.Sqrt(25))
	fmt.Println(math.Pi) //Pi is an exported name from math package.
	fmt.Println(add(4, 5))
	fmt.Println(sub(10, 5))
	e, f := swap("yusuf", "ahmad")
	fmt.Println(e, f)
	fmt.Println(split(100))
	fmt.Println(var1, var2, var3)
	var de int
	ffff := uint64(99)
	gggg := 99.7
	const name = "Yusuf"
	gk := 4 //Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type
	fmt.Println(ab, bc, cd, de)
	fmt.Println(ef, gh, ig, gk)
	//fmt.Println(sqroot, sqroot)
	fmt.Printf(" %v %v %q\n", fff, bbb, sss)
	fmt.Println(ffff)
	fmt.Printf("gggg is a type of %T\n", gggg)

	fmt.Println("My name is \n ", name)
	sum := 1
	//for loop
	for i := 0; i < 3; i++ {
		sum = sum + i
	}
	fmt.Println("Sum is \n ", sum)
	add := 1
	for add < 3 { //Do while in GO
		add += add
	}
	fmt.Println(add)
	// for {    //Forever Loop
	// 	fmt.Println("Yusuf")
	// }

}
