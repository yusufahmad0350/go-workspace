package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

type pair struct {
	row, col int
}

const length = 512

var result [][]int

func main() {
	fmt.Printf("Number of CPUs Total is %d\n", runtime.NumCPU())
	runtime.GOMAXPROCS(4)
	fmt.Printf("GOMAXPROCS is %d\n", runtime.GOMAXPROCS(0))

	mat1 := make([][]int, length) //slice initialize
	for i := range mat1 {
		mat1[i] = make([]int, length)
	}
	mat2 := make([][]int, length)
	for i := range mat2 {
		mat2[i] = make([]int, length)
	}
	resultMat := make([][]int, length)
	for i := range resultMat {
		resultMat[i] = make([]int, length)
	}
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {

			mat1[i][j] = rand.Intn(10)
			mat2[i][j] = rand.Intn(10)
		}
	}
	// fmt.Println(mat1, mat2)
	size := len(mat1)
	pairs := make(chan pair, size*size)
	var wg sync.WaitGroup
	wg.Add(size * size)
	start := time.Now()
	for i := 0; i < size*size; i++ {
		go cannonAlgo(pairs, &mat1, &mat2, &resultMat, &wg) //we dont want copy
	}
	fmt.Println("The no of go routine currently running are : ", runtime.NumGoroutine())
	elapsed := time.Since(start)
	fmt.Print("Time taken to calculate :", elapsed)

	// for i := 0; i < length; i++ {
	// 	for j := 0; j < length; j++ {
	// 		fmt.Print(result[i][j])
	// 		fmt.Print(" ")
	// 	}
	// 	fmt.Println(" ")
	// }
}

func cannonAlgo(pairs chan pair, mat1, mat2, resultMat *[][]int, wg *sync.WaitGroup) {
	n := len(*mat1)
	for {
		pair, ok := <-pairs
		if !ok { //When channel is closed, break out of loop
			break
		}
		for k := 0; k < n; k++ { //Shifting
			i, j := pair.row, pair.col
			if k == 0 { //Each steps forwards
				(*resultMat)[i][j] += (*mat1)[i][(j+i)%n] * (*mat2)[(i+j)%n][j]
			} else {
				(*resultMat)[i][j] += (*mat1)[i][(j+i+k)%n] * (*mat2)[(i+j+k)%n][j]
			}
		}
	}
	wg.Done()
}
