package main

import (
	"fmt"
	"reflect"
	//"runtime"
	"sync"
	"testharness/matrices"
	"time"
)

type Test struct {
	mat1   [][]int
	mat2   [][]int
	result [][]int
}

func main() {
	tests := []Test{
		{mat1: matrices.M0, mat2: matrices.M1, result: matrices.R01},
		{mat1: matrices.M2, mat2: matrices.M3, result: matrices.R23},
		{mat1: matrices.M4, mat2: matrices.M5, result: matrices.R45},
		{mat1: matrices.M6, mat2: matrices.M7, result: matrices.R67}, //512
		{mat1: matrices.M8, mat2: matrices.M9, result: matrices.R89},
		{mat1: matrices.M10, mat2: matrices.M11, result: matrices.R1011}, //1024
	}

	var totalTime int64

	for i, test := range tests {
		var subtestTimeSum int64
		var numIterations int = 10

		testPassed := true

		for subtestNum := 0; subtestNum < numIterations; subtestNum++ {
			start := time.Now()
			result := multiply(test.mat1, test.mat2)
			elapsed := time.Since(start)

			if !reflect.DeepEqual(result, test.result) {
				testPassed = false
				break
			}

			subtestTimeSum += elapsed.Milliseconds()
		}

		fmt.Printf("(%v, %v) x (%v, %v)\n", len(test.mat1), len(test.mat1[0]), len(test.mat2), len(test.mat2[0]))
		if !testPassed {
			fmt.Printf("Error in test %v", i+1)
		} else {
			subtestTimeAvg := subtestTimeSum / int64(numIterations)
			fmt.Printf("Test %v passed in %v ms (average)\n", i+1, subtestTimeAvg)
			totalTime += subtestTimeAvg
		}

		fmt.Println()
	}

	fmt.Printf("\nTotal time: %v ms", totalTime)
}

func multiply(mat1 [][]int, mat2 [][]int) [][]int {
	type pair struct {
		row, col int
	}

	m := len(mat1)    // number of rows the first matrix
	p := len(mat2)    // number of rows the second matrix
	q := len(mat2[0]) // number of columns the second matrix
	resultMat := make([][]int, m)
	for i := 0; i < m; i++ {
		resultMat[i] = make([]int, q)
	}
	pairs := make(chan pair, 50000)
	var wg sync.WaitGroup
	//var mu sync.RWMutex
	wg.Add(1)
	for i := 0; i < 1; i++ {
		go func(pairs chan pair) {
			//wg.Add(1)
			for {
				//wg.Add(1)
				pair, ok := <-pairs //Waiting for pair receive from pairs
				if !ok {
					break
				}
				//mu.Lock()
				wg.Add(1)
				go func() {
					defer wg.Done()

					for k := 0; k < p; k++ {
						resultMat[pair.row][pair.col] += mat1[pair.row][k] * mat2[k][pair.col]

					}
					//mu.Unlock()
				}()
			}

			wg.Done()
		}(pairs)

	}
	for i := 0; i < m; i++ {
		for j := 0; j < q; j++ {
			pairs <- pair{row: i, col: j}
		}
	}
	close(pairs)
	wg.Wait()
	return resultMat
}
