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

//struct

const size = 512

var start time.Time
var resultMat [size][size]int

func main() {
	fmt.Printf("Number of CPUs Total is %d\n", runtime.NumCPU())
	runtime.GOMAXPROCS(4)
	fmt.Printf("GOMAXPROCS is %d\n", runtime.GOMAXPROCS(0))

	pairs := make(chan pair, 50000)
	var wg sync.WaitGroup
	var mat1 [size][size]int
	var mat2 [size][size]int
	for i := 0; i < size; i++ {
		func(i int) {
			for j := 0; j < size; j++ {

				mat1[i][j] = rand.Intn(10)
				mat2[i][j] = rand.Intn(10)
			}
		}(i)
	}
	wg.Add(1)
	for i := 0; i < 1; i++ {
		go Calc(pairs, &mat1, &mat2, &resultMat, &wg) // copy is expensive -pass by reference
	}

	start = time.Now()
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			pairs <- pair{row: i, col: j}
		}
	}
	fmt.Println("The no of go routine currently running are : ", runtime.NumGoroutine())
	//fmt.Println("The first matrix is : ", mat1)
	//fmt.Println("The second Matrix is :", mat2)
	close(pairs)
	wg.Wait()
	elapsed := time.Since(start)
	//fmt.Println("The final Result matrix is ", resultMat) //Print Matrix Results
	fmt.Println("Time taken to calculate :", elapsed.Milliseconds())

}

func Calc(pairs chan pair, mat1, mat2, resultMat *[size][size]int, wg *sync.WaitGroup) {
	for {
		pair, ok := <-pairs //Waiting for pair receive from pairs
		if !ok {
			break
		}
		for i := 0; i < size; i++ {
			resultMat[pair.row][pair.col] += mat1[pair.row][i] * mat2[i][pair.col]
		}
	}
	wg.Done()
}
