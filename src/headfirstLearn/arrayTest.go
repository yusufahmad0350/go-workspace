package main

import (
	"fmt"
)

func main() {
	var myArray [3]int
	myArray[0] = 1
	myArray[1] = 2
	myArray[2]++ //Array 2 has 0 value, incremented the from 0 to 1
	fmt.Println(myArray[2])

	myArray1 := [3]string{"Yusuf", "Danish", "English"}
	fmt.Println(myArray1)
	myArray2 := myArray1[1:]

	fmt.Println(myArray2)
	//Index in array
	index := 1
	fmt.Println(index, myArray1[index])

	for i := 0; i < len(myArray1); i++ {
		fmt.Println(myArray1[i])

	}
	for index, arr := range myArray1 {
		fmt.Println(index, arr)

	}
	for _, arr := range myArray1 {
		fmt.Println(arr)

	}

	//Sum of 3 float numbers
	floatnum := [3]float64{1.4, 1.8, 1.2}
	var sum float64
	for _, fltnum := range floatnum {
		sum += fltnum

	}
	fmt.Println(sum)

	numbers := [6]int{3, 16, -2, 10, 23, 12}
	for i, number := range numbers {
		if number >= 10 && number <= 20 {
			fmt.Println(i, number)
		}
	}
	slice1 := make([]string, 5)
	slice1[0] = "Yusuf"
	slice1[1] = "Danish"
	fmt.Println(slice1)

	slice2 := []string{"abc", "def", "geh"}
	for _, literals := range slice2 {
		fmt.Println(literals)
	}
	slice2 = append(slice2, "mej", "fch")
	fmt.Println(slice2)

}
