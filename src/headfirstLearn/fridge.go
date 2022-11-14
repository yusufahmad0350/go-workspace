package main

import (
	"fmt"
)

func main() {
	var originalCount = 10
	var eatenCount = 4

	fNam := "Yusuf"
	fmt.Println("I started with", originalCount, "apples.")
	fmt.Println("Some Jerk ate", eatenCount, "apples.")
	fmt.Println("There are", originalCount-eatenCount, "apples left")
	fmt.Println(fNam)
}
