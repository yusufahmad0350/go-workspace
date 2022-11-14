package main

import (
	"fmt"
)

func main() {
	var n = []int{1, 39, 2, 9, 7, 54, 11}
	for i := 0; i < len(n); i++ {
		for j := 0; j < len(n)-i-1; j++ {

			if n[j] > n[j+1] {
				n[j], n[j+1] = n[j+1], n[j]
			}

		}

	}
	fmt.Println(n)
}
