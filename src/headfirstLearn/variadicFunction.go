package main

import (
	"fmt"
	"math"
)

func maximum(numbers ...float64) float64 {
	max := math.Inf(-1)
	for _, num := range numbers {
		if num > max {
			max = num

		}

	}
	return max
}

func main() {
	fmt.Println(maximum(10.5, 4.5))
	fmt.Println(maximum(2.3, 3.4, 5.7))
}
