package main

import (
	"fmt"
	"sort"
)

func main() {
	mapMake := make(map[string]int)
	mapMake["yusuf"] = 55
	mapMake["danish"] = 88
	mapMake["qamar"] = 100
	for name, grade := range mapMake {

		if grade > 90 {
			println("You are the topper")
			fmt.Printf("The name is %f and the grade is %v \n", name, grade)
		}
	}
	var namess []string
	for names := range mapMake {
		namess = append(namess, names)

	}
	sort.Strings(namess)

	fmt.Println("The names are", namess)
	mapMake2 := make(map[int]string)
	mapMake2[4] = "Forth"
	mapMake2[5] = "Fifth"
	mapMake2[6] = "Sixth"
	fmt.Println(mapMake2)

	mapMake3 := map[int]string{1: "first", 2: "second", 3: "third"}
	fmt.Println(mapMake3)

	mapMake4 := make(map[string]int) //null map
	mapMake4["a"]++
	mapMake4["a"]++ //incremented twice
	mapMake4["b"]++
	fmt.Println(mapMake4)

	delete(mapMake4, "b")
	fmt.Println(mapMake4)

}
