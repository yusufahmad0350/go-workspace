package main

import (
	"fmt"
	"sync"
)

var matrixA = [][]int{{1, 2}, {3, 4}}
var matrixB = [][]int{{1, 2}, {3, 4}}

var matrixC = make([][]int, 2)

func main() {

	calculateMatrix()

}
func calculateMatrix() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		matrixC[i] = make([]int, 10)
		wg.Add(1)
		func(i int) {

			for j := 0; j < 10; j++ {

				go func(i, j int) {

					for k := 0; k < 10; k++ {
						matrixC[i][j] += matrixA[i][k] * matrixB[k][j]
					}

				}(i, j)
			}
			wg.Done()

		}(i)

	}
	fmt.Print(matrixC)


}
