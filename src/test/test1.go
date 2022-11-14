package main

import (
	"fmt"
	"sync"
)

type result struct {
	name  string
	roll  uint
	marks uint
}

var wg = sync.WaitGroup{}

func main() {
	var v1 = result{"Yusuf", 38, 88}
	fmt.Println(v1)
	var v2 = result{name: "Yusuf"}
	fmt.Println(v2)
	var v3 = result{}
	fmt.Println(v3)
	v1.marks = 66
	fmt.Println(v1)
	var v4 = &result{"Yusuf", 20, 10}
	fmt.Println(v4)
	wg.Add(1)
	name := []string{"Yusuf", "Dsnish", "qamar"}
	var slice2 = name[2:3]
	fmt.Printf("the slice %v is ", slice2)
	func(message string) {
		fmt.Println(message)
	}("Hi, it's me!")
	a := make([]int, 3, 5)
	a=append(a,30)
	//a = a[0:2]
	fmt.Println(a)
	//printSlice(a)

}
